package user

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/eure/si2018-second-half-2/libs/token"
	userlib "github.com/eure/si2018-second-half-2/libs/user"
	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func GetUsers(p si.GetUsersParams) middleware.Responder {
	// バリデーション
	limit := p.Limit
	offset := p.Offset
	t := p.Token
	if res := ValidateGetUsers(limit, offset, t); res != nil {
		return res
	}

	s := repositories.NewSession()

	// Meチェック
	me, err := token.GetUserByToken(s, t)
	if err != nil {
		return si.NewGetUsersInternalServerError().WithPayload(
			&si.GetUsersInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetUsersUnauthorized().WithPayload(
			&si.GetUsersUnauthorizedBody{
				Code:    "401",
				Message: "Token Is Invalid :: 無効なトークン",
			})
	}

	// 探す画面では除外したいIDを取得 (likeした・された一覧をとる)
	likeRp := repositories.NewUserLikeRepository(s)
	likeIDs, err := likeRp.FindLikeAll(me.ID)
	if err != nil {
		return si.NewGetUsersInternalServerError().WithPayload(
			&si.GetUsersInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: いいね取得に失敗しました",
			})
	}

	// ユーザーの取得
	r := repositories.NewUserRepository(s)
	users, err := r.FindWithCondition(int(limit), int(offset), me.GetOppositeGender(), likeIDs)
	if err != nil {
		return si.NewGetUsersInternalServerError().WithPayload(
			&si.GetUsersInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: ユーザーの取得に失敗しました",
			})
	}

	userList, err := userlib.SetUsersImage(s, users)
	if err != nil {
		return si.NewGetUsersInternalServerError().WithPayload(
			&si.GetUsersInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: 画像の取得に失敗しました",
			})
	}

	eUsers := entities.Users(userList)
	sUsers := eUsers.Build()
	return si.NewGetUsersOK().WithPayload(sUsers)
}

func GetProfileByUserID(p si.GetProfileByUserIDParams) middleware.Responder {
	// バリデーション
	t := p.Token
	if res := ValidateGetProfileByUserID(t); res != nil {
		return res
	}

	s := repositories.NewSession()

	// Meの取得
	me, err := token.GetUserByToken(s, t)
	if err != nil {
		return si.NewGetProfileByUserIDInternalServerError().WithPayload(
			&si.GetProfileByUserIDInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetProfileByUserIDUnauthorized().WithPayload(
			&si.GetProfileByUserIDUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: 無効なトークン",
			})
	}

	// ユーザーの取得
	r := repositories.NewUserRepository(s)
	ent, err := r.GetByUserID(p.UserID)
	if err != nil {
		return si.NewGetProfileByUserIDInternalServerError().WithPayload(
			&si.GetProfileByUserIDInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: ユーザーの取得に失敗しました",
			})
	}
	if ent == nil {
		return si.NewGetProfileByUserIDNotFound().WithPayload(
			&si.GetProfileByUserIDNotFoundBody{
				Code:    "404",
				Message: "User Not Found :: ユーザーが存在しません",
			})
	}

	user, err := userlib.SetUserImage(s, *ent)
	if err != nil {
		return si.NewGetProfileByUserIDInternalServerError().WithPayload(
			&si.GetProfileByUserIDInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: 画像の取得に失敗しました",
			})
	}

	sEnt := user.Build()
	return si.NewGetProfileByUserIDOK().WithPayload(&sEnt)
}

func PutProfile(p si.PutProfileParams) middleware.Responder {
	// バリデーション
	t := p.Params.Token
	uri := p.Params.ImageURI
	if res := ValidatePutProfile(t, uri); res != nil {
		return res
	}

	s := repositories.NewSession()

	// Meチェック
	me, err := token.GetUserByToken(s, t)
	if err != nil {
		return si.NewPutProfileInternalServerError().WithPayload(
			&si.PutProfileInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewPutProfileUnauthorized().WithPayload(
			&si.PutProfileUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: 無効なトークン",
			})
	}

	// 他人のプロフィールは書き換えられない
	if p.UserID != me.ID {
		return si.NewPutProfileForbidden().WithPayload(
			&si.PutProfileForbiddenBody{
				Code:    "403",
				Message: "Forbidden :: 更新できるのは自分のプロフィールのみです",
			})
	}

	// プロフィールのアップデート
	r := repositories.NewUserRepository(s)
	ent := userlib.BuildUserEntityByModel(me.ID, p.Params)
	err = r.Update(&ent)
	if err != nil {
		return si.NewPutProfileInternalServerError().WithPayload(
			&si.PutProfileInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: プロフィールの更新に失敗しました",
			})
	}

	// 更新されたuser entityを取得してまるっと返す
	user, err := r.GetByUserID(me.ID)
	if err != nil {
		return si.NewPutProfileInternalServerError().WithPayload(
			&si.PutProfileInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: プロフィールの取得に失敗しました",
			})
	}
	rp := repositories.NewUserImageRepository(s)
	userImage, err := rp.GetByUserID(user.ID)
	if err != nil {
		return si.NewPutProfileInternalServerError().WithPayload(
			&si.PutProfileInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: 画像の取得に失敗しました",
			})
	}
	user.ImageURI = userImage.Path

	sEnt := user.Build()
	return si.NewPutProfileOK().WithPayload(&sEnt)
}
