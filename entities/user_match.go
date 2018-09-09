package entities

import (
	"github.com/eure/si2018-second-half-2/models"
	"github.com/go-openapi/strfmt"
)

type UserMatch struct {
	UserID    int64           `xorm:"user_id"`
	PartnerID int64           `xorm:"partner_id"`
	CreatedAt strfmt.DateTime `xorm:"created_at"`
	UpdatedAt strfmt.DateTime `xorm:"updated_at"`
}

func (u UserMatch) Build() models.UserMatch {
	return models.UserMatch{
		UserID:    u.UserID,
		PartnerID: u.PartnerID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type UserMatches []UserMatch

func (uu *UserMatches) Build() []*models.UserMatch {
	var sUsers []*models.UserMatch

	for _, u := range *uu {
		swaggerUser := u.Build()
		sUsers = append(sUsers, &swaggerUser)
	}
	return sUsers
}
