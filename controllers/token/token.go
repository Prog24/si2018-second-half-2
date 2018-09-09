package token

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func GetTokenByUserID(p si.GetTokenByUserIDParams) middleware.Responder {
	s := repositories.NewSession()
	r := repositories.NewUserTokenRepository(s)
	ent, err := r.GetByUserID(p.UserID)
	if err != nil {
		return si.NewGetTokenByUserIDInternalServerError().WithPayload(
			&si.GetTokenByUserIDInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: 取得に失敗しました",
			})
	}
	if ent == nil {
		return si.NewGetTokenByUserIDNotFound().WithPayload(
			&si.GetTokenByUserIDNotFoundBody{
				Code:    "404",
				Message: "User Token Not Found :: トークンが存在しません",
			})
	}

	sEnt := ent.Build()
	return si.NewGetTokenByUserIDOK().WithPayload(&sEnt)
}
