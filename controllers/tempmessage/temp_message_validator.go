package tempmessage

import (
	"github.com/eure/si2018-second-half-2/repositories"
	"github.com/go-openapi/runtime/middleware"

	t "github.com/eure/si2018-second-half-2/libs/token"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func ValidatePostTempMessage(s *repositories.Session, userID int64, msg, token string) middleware.Responder {
	if res := validateTempMessage(s, msg); res != nil {
		return res
	}
	if res := validatePostTempMessageToken(s, token); res != nil {
		return res
	}
	return nil
}

func validateTempMessage(s *repositories.Session, msg string) middleware.Responder {
	if len(msg) == 0 {
		return si.NewPostTempMessageBadRequest().WithPayload(
			&si.PostTempMessageBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: パラメータが不足です",
			})
	}
	return nil
}

func validatePostTempMessageToken(s *repositories.Session, token string) middleware.Responder {
	me, err := t.GetUserByToken(s, token)
	if err != nil {
		return si.NewPostTempMessageInternalServerError().WithPayload(
			&si.PostTempMessageInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewPostTempMessageUnauthorized().WithPayload(
			&si.PostTempMessageUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}
	return nil
}

func ValidateGetTempMessagesParams(s *repositories.Session, token string) middleware.Responder {
	if res := validateGetTempMessagesToken(s, token); res != nil {
		return res
	}
	return nil
}

func validateGetTempMessagesToken(s *repositories.Session, token string) middleware.Responder {
	me, err := t.GetUserByToken(s, token)
	if err != nil {
		return si.NewGetTempMessagesInternalServerError().WithPayload(
			&si.GetTempMessagesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetTempMessagesUnauthorized().WithPayload(
			&si.GetTempMessagesUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}
	return nil
}
