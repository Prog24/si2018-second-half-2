package usermatch

import (
	"github.com/go-openapi/runtime/middleware"

	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

type Validator interface {
	Validate() middleware.Responder
}

type GetValidator struct {
	token  string
	limit  int64
	offset int64
}

func NewGetValidator(t string, l, o int64) Validator {
	return GetValidator{
		token:  t,
		limit:  l,
		offset: o,
	}
}

func (v GetValidator) Validate() middleware.Responder {
	if v.limit == 0 {
		return si.NewGetMatchesBadRequest().WithPayload(
			&si.GetMatchesBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: limitを指定してください",
			})
	}

	if len(v.token) == 0 {
		return si.NewGetMatchesUnauthorized().WithPayload(
			&si.GetMatchesUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: tokenを指定してください",
			})
	}

	return nil
}
