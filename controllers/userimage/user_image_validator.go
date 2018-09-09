package userimage

import (
	"github.com/go-openapi/runtime/middleware"

	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

type Validator interface {
	Validate() middleware.Responder
}

type PostValidator struct {
	token       string
	imageBase64 []byte
}

func NewPostValidator(t string, i []byte) Validator {
	return PostValidator{
		token:       t,
		imageBase64: i,
	}
}

func (v PostValidator) Validate() middleware.Responder {
	if len(v.imageBase64) == 0 {
		return si.NewPostImagesBadRequest().WithPayload(
			&si.PostImagesBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: 画像データが必須です",
			})
	}

	if len(v.token) == 0 {
		return si.NewPostImagesUnauthorized().WithPayload(
			&si.PostImagesUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: Tokenを指定してください",
			})
	}

	return nil
}
