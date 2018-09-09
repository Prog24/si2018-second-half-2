package message

import (
	"github.com/eure/si2018-second-half-2/repositories"
	"github.com/go-openapi/runtime/middleware"

	t "github.com/eure/si2018-second-half-2/libs/token"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func ValidatePostMessage(s *repositories.Session, userID int64, msg, token string) middleware.Responder {
	if res := validateMessage(s, msg); res != nil {
		return res
	}
	if res := validatePostMessageToken(s, token); res != nil {
		return res
	}
	return nil
}

func validateMessage(s *repositories.Session, msg string) middleware.Responder {
	if len(msg) == 0 {
		return si.NewPostMessageBadRequest().WithPayload(
			&si.PostMessageBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: パラメータが不足です",
			})
	}
	return nil
}

func validatePostMessageToken(s *repositories.Session, token string) middleware.Responder {
	me, err := t.GetUserByToken(s, token)
	if err != nil {
		return si.NewPostMessageInternalServerError().WithPayload(
			&si.PostMessageInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewPostMessageUnauthorized().WithPayload(
			&si.PostMessageUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}
	return nil
}

func ValidateGetMessagesParams(s *repositories.Session, token string) middleware.Responder {
	if res := validateGetMessagesToken(s, token); res != nil {
		return res
	}
	return nil
}

func validateGetMessagesToken(s *repositories.Session, token string) middleware.Responder {
	me, err := t.GetUserByToken(s, token)
	if err != nil {
		return si.NewGetMessagesInternalServerError().WithPayload(
			&si.GetMessagesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewGetMessagesUnauthorized().WithPayload(
			&si.GetMessagesUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: トークンが無効です",
			})
	}
	return nil
}
