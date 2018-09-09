package user

import (
	"github.com/go-openapi/runtime/middleware"

	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func ValidateGetUsers(limit, offset int64, t string) middleware.Responder {
	if limit == 0 {
		return si.NewGetUsersBadRequest().WithPayload(
			&si.GetUsersBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: limitを指定してください",
			})
	}

	if len(t) == 0 {
		return si.NewGetUsersUnauthorized().WithPayload(
			&si.GetUsersUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: Tokenを指定してください",
			})
	}

	return nil
}

func ValidateGetProfileByUserID(t string) middleware.Responder {
	if len(t) == 0 {
		return si.NewGetProfileByUserIDUnauthorized().WithPayload(
			&si.GetProfileByUserIDUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: Tokenを指定してください",
			})
	}

	return nil
}

func ValidatePutProfile(t, uri string) middleware.Responder {
	if len(t) == 0 {
		return si.NewPutProfileUnauthorized().WithPayload(
			&si.PutProfileUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: Tokenを指定してください",
			})
	}

	if len(uri) != 0 {
		return si.NewPutProfileBadRequest().WithPayload(
			&si.PutProfileBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: 画像のアップデートは POST /images/ で行ってください",
			})
	}

	return nil
}
