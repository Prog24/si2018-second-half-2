package entities

import (
	"github.com/eure/si2018-second-half-2/models"
	"github.com/go-openapi/strfmt"
)

type UserWaitTempMatch struct {
	UserID     int64           `xorm:"user_id"`
	Gender     string          `xorm:"gender"`
	IsMatched  bool            `xorm:"is_matched"`
	IsCanceled bool            `xorm:"is_canceled"`
	CreatedAt  strfmt.DateTime `xorm:"created_at"`
	UpdatedAt  strfmt.DateTime `xorm:"updated_at"`
}

func (u UserWaitTempMatch) Build() models.UserWaitTempMatch {
	return models.UserWaitTempMatch{
		UserID:     u.UserID,
		Gender:     u.Gender,
		IsMatched:  u.IsMatched,
		IsCanceled: u.IsCanceled,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

type UserWaitTempMatches []UserWaitTempMatch

func (uu *UserWaitTempMatches) Build() []*models.UserWaitTempMatch {
	var sUsers []*models.UserWaitTempMatch

	for _, u := range *uu {
		swaggerUser := u.Build()
		sUsers = append(sUsers, &swaggerUser)
	}
	return sUsers
}
