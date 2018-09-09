package userimage

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/eure/si2018-second-half-2/libs/token"
	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func PostImage(p si.PostImagesParams) middleware.Responder {
	t := p.Params.Token
	v := NewPostValidator(t, p.Params.Image)
	if res := v.Validate(); res != nil {
		return res
	}

	s := repositories.NewSession()

	me, err := token.GetUserByToken(s, t)
	if err != nil {
		return si.NewPostImagesInternalServerError().WithPayload(
			&si.PostImagesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Meの取得に失敗しました",
			})
	}
	if me == nil {
		return si.NewPostImagesUnauthorized().WithPayload(
			&si.PostImagesUnauthorizedBody{
				Code:    "401",
				Message: "Unauthorized :: Tokenが無効です",
			})
	}

	r := repositories.NewUserImageRepository(s)
	now := strfmt.DateTime(time.Now())

	hostID := os.Getenv("HOST_ID")         // 000〜020
	assetsPath := os.Getenv("ASSETS_PATH") // $GOPATH/src/github.com/eure/si2018-second-half-2/assets/

	timeStamp := fmt.Sprintf("%d", time.Now().Unix())
	meIDStr := fmt.Sprintf("%d", me.ID)

	reader := bytes.NewReader(p.Params.Image)
	_, format, err := image.DecodeConfig(reader)
	if err != nil {
		// 画像フォーマットではない場合はエラーが発生する
		return si.NewPostImagesBadRequest().WithPayload(
			&si.PostImagesBadRequestBody{
				Code:    "400",
				Message: "Bad Request :: ファイルのフォーマットが不適切です",
			})
	}
	imgType := format
	imgName := hostID + "_" + meIDStr + "_" + timeStamp + "." + imgType

	imgPath := assetsPath + imgName // サーバー上での置き場所

	file, err := os.Create(imgPath)
	if err != nil {
		return si.NewPostImagesInternalServerError().WithPayload(
			&si.PostImagesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: 画像のアップロードに失敗しました",
			})
	}
	defer file.Close()

	file.Write(p.Params.Image)

	// https://si-2018-000.eure.jp/assets/hogehoge.jpg
	imgURI := os.Getenv("ASSETS_BASE_URI") + imgName
	userImage := entities.UserImage{
		UserID:    me.ID,
		Path:      imgURI,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// CreateOrUpdateしている (ImageはUserユニークなレコード)
	image, err := r.GetByUserID(me.ID)
	if err != nil {
		return si.NewPostImagesInternalServerError().WithPayload(
			&si.PostImagesInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: 画像取得に失敗しました",
			})
	}
	if image == nil {
		err = r.Create(userImage)
		if err != nil {
			return si.NewPostImagesInternalServerError().WithPayload(
				&si.PostImagesInternalServerErrorBody{
					Code:    "500",
					Message: "Internal Server Error :: 画像作成に失敗しました",
				})
		}
	} else {
		err = r.Update(userImage)
		if err != nil {
			return si.NewPostImagesInternalServerError().WithPayload(
				&si.PostImagesInternalServerErrorBody{
					Code:    "500",
					Message: "Internal Server Error :: 画像更新に失敗しました",
				})
		}
	}

	payload := si.PostImagesOKBody{
		ImageURI: strfmt.URI(imgURI),
	}
	return si.NewPostImagesOK().WithPayload(&payload)
}
