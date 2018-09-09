package userlike

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/eure/si2018-second-half-2/libs/token"
	userlib "github.com/eure/si2018-second-half-2/libs/user"
	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func GetLikes(p si.GetLikesParams) middleware.Responder {
	// バリデーション
	t := p.Token
	limit := p.Limit
	offset := p.Offset
	v := NewGetValidator(t, limit, offset)
	if res := v.Validate(); res != nil {
		return res
	}

	s := repositories.NewSession()

	// meチェック
	me, err := token.GetUserByToken(s, t)
	if err != nil {
		return si.NewGetLikesInternalServerError().WithPayload(
			&si.GetLikesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetLikesUnauthorized().WithPayload(
			&si.GetLikesUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}

	// LikeのレスポンスにはMatch済みのお相手を返さないので not in するためにmatchIDsを取得する
	mr := repositories.NewUserMatchRepository(s)
	matchIDs, err := mr.FindAllByUserID(me.ID)

	// GotLikeをlimit offsetで取得
	r := repositories.NewUserLikeRepository(s)
	likes, err := r.FindGotLikeWithLimitOffset(me.ID, int(limit), int(offset), matchIDs)
	if err != nil {
		return si.NewGetLikesInternalServerError().WithPayload(
			&si.GetLikesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: likeの取得に失敗",
			})
	}

	// 取得したlikesをrangeで回す
	// keyにお相手のUID, valueがLikeUserResponseのlikeUserMapを作る
	// ( LikeUserResponse.liked_at はここで入れる )
	likeUserIDList := make([]int64, len(likes))
	likeUserMap := make(map[int64]entities.LikeUserResponse, len(likes))
	for i, l := range likes {
		if me.ID == l.UserID {
			likeUserIDList[i] = l.PartnerID
			likeUserMap[l.PartnerID] = entities.LikeUserResponse{
				LikedAt: l.CreatedAt,
			}
			continue
		}
		likeUserIDList[i] = l.UserID
		likeUserMap[l.UserID] = entities.LikeUserResponse{
			LikedAt: l.CreatedAt,
		}
	}

	// ユーザーを取得
	ur := repositories.NewUserRepository(s)
	users, err := ur.FindByIDs(likeUserIDList)
	if err != nil {
		return si.NewGetLikesInternalServerError().WithPayload(
			&si.GetLikesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Userの取得に失敗",
			})
	}

	// UserImage.URIをUser.ImageURIに付与
	users, err = userlib.SetUsersImage(s, users)
	if err != nil {
		return si.NewGetLikesInternalServerError().WithPayload(
			&si.GetLikesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: UserImageの取得に失敗",
			})
	}

	// Like順に, UserデータをlikeUserMapに入れていく
	for _, u := range users {
		respUser := likeUserMap[u.ID]
		respUser.ApplyUser(u)
		likeUserMap[u.ID] = respUser
	}

	// MapをSliceに変換
	sortedRespList := make([]entities.LikeUserResponse, len(likeUserIDList))
	for i, m := range likeUserIDList {
		sortedRespList[i] = likeUserMap[m]
	}

	// SwaggerのModelに変換してreturnする
	responses := entities.LikeUserResponses(sortedRespList)
	swaggerResponses := responses.Build()
	return si.NewGetLikesOK().WithPayload(swaggerResponses)
}

func PostLike(p si.PostLikeParams) middleware.Responder {
	// リクエストパラメータのバリデーション
	t := p.Params.Token
	v := NewPostValidator(t, p.UserID)
	if resp := v.Validate(); resp != nil {
		return resp
	}

	s := repositories.NewSession()

	// meチェック
	me, err := token.GetUserByToken(s, t)
	if err != nil {
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewPostLikeUnauthorized().WithPayload(
			&si.PostLikeUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}

	partnerID := p.UserID
	ur := repositories.NewUserRepository(s)
	partner, err := ur.GetByUserID(partnerID)
	if err != nil {
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Partnerの取得に失敗しました",
			})
	}
	if partner == nil {
		return si.NewPostLikeBadRequest().WithPayload(
			&si.PostLikeBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: お相手が存在しません",
			})
	}
	if me.Gender == partner.Gender {
		return si.NewPostLikeBadRequest().WithPayload(
			&si.PostLikeBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: 同性へのいいねはできません",
			})
	}

	// Like (Me -> Partner) の存在チェック
	// 既に送ってたら400で返す.
	lr := repositories.NewUserLikeRepository(s)
	sendLike, err := lr.GetLikeBySenderIDReceiverID(me.ID, partnerID)
	if err != nil {
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Likeの取得に失敗しました",
			})
	}
	if sendLike != nil {
		return si.NewPostLikeBadRequest().WithPayload(
			&si.PostLikeBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: 既にLike送信済みです",
			})
	}

	repositories.TransactionBegin(s)
	// Like (Me -> Partner) レコードのInsert
	now := strfmt.DateTime(time.Now())
	err = lr.Create(
		entities.UserLike{
			UserID:    me.ID,
			PartnerID: partnerID,
			CreatedAt: now,
			UpdatedAt: now,
		})
	if err != nil {
		repositories.TransactionRollBack(s)
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: LikeのInsertに失敗しました",
			})
	}

	// Like (Partner -> Me) の存在チェック
	// 向こうからまだLikeが来てなければ, LikeをInsertするだけでreturnしてよい
	getLike, err := lr.GetLikeBySenderIDReceiverID(partnerID, me.ID)
	if getLike == nil {
		repositories.TransactionCommit(s)
		return si.NewPostLikeOK().WithPayload(
			&si.PostLikeOKBody{
				Code:    "200",
				Message: "OK :: Likeが送信されました",
			})
	}

	// 向こうからLikeがきていた場合,Like送信して,さらにMatchもさせる
	mr := repositories.NewUserMatchRepository(s)
	err = mr.Create(
		entities.UserMatch{
			UserID:    me.ID,
			PartnerID: partnerID,
			CreatedAt: now,
			UpdatedAt: now,
		})
	if err != nil {
		repositories.TransactionRollBack(s)
		return si.NewPostLikeInternalServerError().WithPayload(
			&si.PostLikeInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: MatchのInsertに失敗しました",
			})
	}
	repositories.TransactionCommit(s)

	return si.NewPostLikeOK().WithPayload(
		&si.PostLikeOKBody{
			Code:    "200",
			Message: "OK :: Likeを送信し,お相手とMatchしました",
		})
}
