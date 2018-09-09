package message

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/eure/si2018-second-half-2/libs/token"
	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func PostMessage(p si.PostMessageParams) middleware.Responder {
	s := repositories.NewSession()

	// バリデーション
	msg := p.Params.Message
	t := p.Params.Token
	partnerID := p.UserID
	if res := ValidatePostMessage(s, partnerID, msg, t); res != nil {
		return res
	}

	me, err := token.GetUserByToken(s, t)
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

	userRepo := repositories.NewUserRepository(s)
	partner, err := userRepo.GetByUserID(partnerID)
	if err != nil {
		return si.NewPostMessageInternalServerError().WithPayload(
			&si.PostMessageInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: お相手の取得に失敗しました",
			})
	}
	if partner == nil {
		return si.NewPostMessageBadRequest().WithPayload(
			&si.PostMessageBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: お相手が存在しません",
			})
	}
	if me.Gender == partner.Gender {
		return si.NewPostMessageBadRequest().WithPayload(
			&si.PostMessageBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: 同性へのメッセージはできません",
			})
	}

	// マッチの判定
	matchRepo := repositories.NewUserMatchRepository(s)
	match, err := matchRepo.Get(me.ID, partnerID)
	if err != nil {
		return si.NewPostMessageInternalServerError().WithPayload(
			&si.PostMessageInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: マッチの取得に失敗しました",
			})
	}
	if match == nil {
		return si.NewPostMessageBadRequest().WithPayload(
			&si.PostMessageBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: お相手とまだマッチしていません",
			})
	}

	// Messageレコードの作成
	msgRepo := repositories.NewUserMessageRepository(s)

	now := strfmt.DateTime(time.Now())
	ent := entities.UserMessage{
		UserID:    me.ID,
		PartnerID: p.UserID,
		Message:   msg,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = msgRepo.Create(ent)
	if err != nil {
		// FIXME :: 1秒以内に送ると Error 1062: Duplicate entry '2-1112-2018-09-01 20:37:08' for key 'PRIMARY'
		return si.NewPostMessageInternalServerError().WithPayload(
			&si.PostMessageInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: メッセージの送信に失敗しました",
			})
	}

	return si.NewPostMessageOK().WithPayload(
		&si.PostMessageOKBody{
			Code:    "200",
			Message: "OK :: メッセージを送信しました",
		})
}

func GetMessages(p si.GetMessagesParams) middleware.Responder {
	s := repositories.NewSession()

	// バリデーション
	latest := p.Latest
	oldest := p.Oldest
	limit := p.Limit
	t := p.Token
	if res := ValidateGetMessagesParams(s, t); res != nil {
		return res
	}

	// Meの取得
	me, err := token.GetUserByToken(s, t)
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

	r := repositories.NewUserMessageRepository(s)
	msgs, err := r.GetMessages(me.ID, p.UserID, int(limit), latest, oldest)
	if err != nil {
		return si.NewGetMessagesInternalServerError().WithPayload(
			&si.GetMessagesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: メッセージの取得に失敗しました",
			})
	}

	eMsgs := entities.UserMessages(msgs)
	sMsgs := eMsgs.Build()
	return si.NewGetMessagesOK().WithPayload(sMsgs)
}
