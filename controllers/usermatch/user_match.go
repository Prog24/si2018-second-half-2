package usermatch

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/eure/si2018-second-half-2/entities"
	tokenlib "github.com/eure/si2018-second-half-2/libs/token"
	userlib "github.com/eure/si2018-second-half-2/libs/user"
	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func GetMatches(p si.GetMatchesParams) middleware.Responder {
	// バリデーション
	t := p.Token
	limit := p.Limit
	offset := p.Offset
	v := NewGetValidator(t, limit, offset)
	if resp := v.Validate(); resp != nil {
		return resp
	}

	s := repositories.NewSession()

	// meチェック
	me, err := tokenlib.GetUserByToken(s, t)
	if err != nil {
		return si.NewGetMatchesInternalServerError().WithPayload(
			&si.GetMatchesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetMatchesUnauthorized().WithPayload(
			&si.GetMatchesUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}

	// Matchをlimit offsetで取得
	r := repositories.NewUserMatchRepository(s)
	matches, err := r.FindByUserIDWithLimitOffset(me.ID, int(limit), int(offset))
	if err != nil {
		return si.NewGetMatchesInternalServerError().WithPayload(
			&si.GetMatchesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: matchの取得に失敗しました",
			})
	}

	// 取得したmatchesをrangeで回す
	// keyにお相手のUID, valueがLikeUserResponseのlikeUserMapを作る
	// ( MatchUserResponse.matched_at はここで入れる )
	matchUserIDList := make([]int64, len(matches))
	matchUserMap := make(map[int64]entities.MatchUserResponse, len(matches))
	for i, m := range matches {
		// Match.UserIDがいいねを先に送った人 / Match.PartnerIDがいいねを返した人.
		// どちらにもme.IDが入っている可能性がある.
		if me.ID == m.UserID {
			matchUserIDList[i] = m.PartnerID
			matchUserMap[m.PartnerID] = entities.MatchUserResponse{
				MatchedAt: m.CreatedAt,
			}
			continue
		}
		matchUserIDList[i] = m.UserID
		matchUserMap[m.UserID] = entities.MatchUserResponse{
			MatchedAt: m.CreatedAt,
		}
	}

	// Userを取得
	ur := repositories.NewUserRepository(s)
	users, err := ur.FindByIDs(matchUserIDList)
	if err != nil {
		return si.NewGetMatchesInternalServerError().WithPayload(
			&si.GetMatchesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Userの取得に失敗しました",
			})
	}

	// UserImage.URIをUser.ImageURIに付与
	users, err = userlib.SetUsersImage(s, users)
	if err != nil {
		return si.NewGetMatchesInternalServerError().WithPayload(
			&si.GetMatchesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: UserImageの取得に失敗しました",
			})
	}

	// Match順に, UserデータをmatchUserMapに入れていく
	for _, u := range users {
		respUser := matchUserMap[u.ID]
		respUser.ApplyUser(u)
		matchUserMap[u.ID] = respUser
	}

	// MapをSliceに変換
	sortedRespList := make([]entities.MatchUserResponse, len(matches))
	for i, m := range matchUserIDList {
		sortedRespList[i] = matchUserMap[m]
	}

	// SwaggerのModelに変換してreturnする
	responses := entities.MatchUserResponses(sortedRespList)
	swaggerResponses := responses.Build()
	return si.NewGetMatchesOK().WithPayload(swaggerResponses)
}
