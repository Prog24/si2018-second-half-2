package usertempmatch

import (
	"github.com/eure/si2018-second-half-2/repositories"
	"github.com/go-openapi/runtime/middleware"

	t "github.com/eure/si2018-second-half-2/libs/token"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func ValidatePostTempMatch(s *repositories.Session, token string) middleware.Responder {
	if res := validatePostTempMatchToken(s, token); res != nil {
		return res
	}
	return nil
}

func validatePostTempMatchToken(s *repositories.Session, token string) middleware.Responder {
	me, err := t.GetUserByToken(s, token)
	if err != nil {
		return si.NewPostTempMatchInternalServerError().WithPayload(
			&si.PostTempMatchInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewPostTempMatchUnauthorized().WithPayload(
			&si.PostTempMatchUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}
	return nil
}

func ValidateGetTempMatch(s *repositories.Session, token string) middleware.Responder {
	if res := validateGetTempMatchToken(s, token); res != nil {
		return res
	}
	return nil
}

func validateGetTempMatchToken(s *repositories.Session, token string) middleware.Responder {
	me, err := t.GetUserByToken(s, token)
	if err != nil {
		return si.NewGetTempMatchInternalServerError().WithPayload(
			&si.GetTempMatchInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetTempMatchUnauthorized().WithPayload(
			&si.GetTempMatchUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}
	return nil
}

func ValidatePutTempMatch(s *repositories.Session, token string) middleware.Responder {
	if res := validatePutTempMatchToken(s, token); res != nil {
		return res
	}
	return nil
}

func validatePutTempMatchToken(s *repositories.Session, token string) middleware.Responder {
	me, err := t.GetUserByToken(s, token)
	if err != nil {
		return si.NewPutTempMatchInternalServerError().WithPayload(
			&si.PutTempMatchInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewPutTempMatchUnauthorized().WithPayload(
			&si.PutTempMatchUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}
	return nil
}
