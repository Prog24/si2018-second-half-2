package entities

import (
	"github.com/eure/si2018-second-half-2/models"
	"github.com/go-openapi/strfmt"
)

type UserLike struct {
	UserID    int64           `xorm:"user_id"`
	PartnerID int64           `xorm:"partner_id"`
	CreatedAt strfmt.DateTime `xorm:"created_at"`
	UpdatedAt strfmt.DateTime `xorm:"updated_at"`
}

func (u UserLike) Build() models.UserLike {
	return models.UserLike{
		UserID:    u.UserID,
		PartnerID: u.PartnerID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type UserLikes []UserLike

func (users *UserLikes) Build() []*models.UserLike {
	var sUsers []*models.UserLike

	for _, u := range *users {
		swaggerUser := u.Build()
		sUsers = append(sUsers, &swaggerUser)
	}
	return sUsers
}
