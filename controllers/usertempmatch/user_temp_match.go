package usertempmatch

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-2/entities"
	tokenlib "github.com/eure/si2018-second-half-2/libs/token"
	// userlib "github.com/eure/si2018-second-half-2/libs/user"
	"github.com/eure/si2018-second-half-2/repositories"
	si "github.com/eure/si2018-second-half-2/restapi/summerintern"
)

func GetTempMatch(p si.GetTempMatchParams) middleware.Responder {
	s := repositories.NewSession()

	// Validation
	t := p.Token
	if res := ValidateGetTempMatch(s, t); res != nil {
		return res
	}

	//

	return si.NewGetTempMatchOK()
}

func PostTempMatch(p si.PostTempMatchParams) middleware.Responder {
	s := repositories.NewSession()

	// Validation
	t := p.Token
	if res := ValidatePostTempMatch(s, t); res != nil {
		return res
	}

	// Get me
	me, err := tokenlib.GetUserByToken(s, t)
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
				Message: "Unauthorized :: Tokenが無効です",
			})
	}

	// Check if you are active
	/* ---------------
	*
	* include validation
	*
	 */

	// Create UserWaitTempMatch entities for me
	now := strfmt.DateTime(time.Now())
	waitTempMatchEnt := entities.UserWaitTempMatch{
		UserID:     me.ID,
		Gender:     me.Gender,
		IsMatched:  false,
		IsCanceled: false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	r := repositories.NewUserWaitTempMatchRepository(s)
	err = r.Create(waitTempMatchEnt)
	if err != nil {
		return si.NewPostTempMatchInternalServerError().WithPayload(
			&si.PostTempMatchInternalServerErrorBody{
				Code:    "500",
				Message: "Internal Server Error :: Failed to wait temp match",
			})
	}

	// Search suited user for me
	/* ----------
	*
	*
	*
	 */

	return si.NewPostTempMatchOK()
}

func PutTempMatch(p si.PutTempMatchParams) middleware.Responder {
	return si.NewPutTempMatchOK()
}
