package tempmessage

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/eure/si2018-second-half-2/libs/token"
	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func PostTempMessage(p si.PostTempMessageParams) middleware.Responder {
	s := repositories.NewSession()

	// Validation
	msg := p.Params.Message
	t := p.Params.Token
	partnerID := p.UserID
	if res := ValidatePostTempMessage(s, partnerID, msg, t); res != nil {
		return res
	}

	me, err := token.GetUserByToken(s, t)
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

	userRepo := repositories.NewUserRepository(s)
	partner, err := userRepo.GetByUserID(partnerID)
	if err != nil {
		return si.NewPostTempMessageInternalServerError().WithPayload(
			&si.PostTempMessageInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: お相手の取得に失敗しました",
			})
	}
	if partner == nil {
		return si.NewPostTempMessageBadRequest().WithPayload(
			&si.PostTempMessageBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: お相手が存在しません",
			})
	}
	if me.Gender == partner.Gender {
		return si.NewPostTempMessageBadRequest().WithPayload(
			&si.PostTempMessageBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: 同性へのメッセージはできません",
			})
	}

	// Check I matched the partner
	tempmatchRepo := repositories.NewUserTempMatchRepository(s)
	tempmatch, err := tempmatchRepo.Get(me.ID, partner.ID)
	if err != nil {
		return si.NewPostTempMessageInternalServerError().WithPayload(
			&si.PostTempMessageInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: マッチの取得に失敗しました",
			})
	}
	if tempmatch == nil {
		return si.NewPostTempMessageBadRequest().WithPayload(
			&si.PostTempMessageBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: お相手とまだマッチしていません",
			})
	}

	// Create message records
	tempmsgRepo := repositories.NewUserTempMessageRepository(s)

	now := strfmt.DateTime(time.Now())
	ent := entities.UserTempMessage{
		UserID:    me.ID,
		PartnerID: p.UserID,
		Message:   msg,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = tempmsgRepo.Create(ent)
	if err != nil {
		return si.NewPostTempMessageInternalServerError().WithPayload(
			&si.PostTempMessageInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: メッセージの送信に失敗しました",
			})
	}

	return si.NewPostTempMessageOK().WithPayload(
		&si.PostTempMessageOKBody{
			Code:    "200",
			Message: "OK :: メッセージを送信しました",
		})
}

func GetTempMessages(p si.GetTempMessagesParams) middleware.Responder {
	s := repositories.NewSession()

	// Validation
	latest := p.Latest
	oldest := p.Oldest
	limit := p.Limit
	t := p.Token
	if res := ValidateGetTempMessagesParams(s, t); res != nil {
		return res
	}

	// Get me
	me, err := token.GetUserByToken(s, t)
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

	r := repositories.NewUserTempMessageRepository(s)
	msgs, err := r.GetMessages(me.ID, p.UserID, int(limit), latest, oldest)
	if err != nil {
		return si.NewGetTempMessagesInternalServerError().WithPayload(
			&si.GetTempMessagesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: メッセージの取得に失敗しました",
			})
	}

	eMsgs := entities.UserTempMessages(msgs)
	sMsgs := eMsgs.Build()
	return si.NewGetTempMessagesOK().WithPayload(sMsgs)
}
