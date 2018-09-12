package entities

import (
	"github.com/eure/si2018-second-half-2/models"
	"github.com/go-openapi/strfmt"
)

type UserTempMatch struct {
	UserID    int64           `xorm:"user_id"`
	PartnerID int64           `xorm:"partner_id"`
	CreatedAt strfmt.DateTime `xorm:"created_at"`
	UpdatedAt strfmt.DateTime `xorm:"updated_at"`
}

func (u UserTempMatch) Build() models.UserTempMatch {
	return models.UserTempMatch{
		UserID:    u.UserID,
		PartnerID: u.PartnerID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type UserTempMatches []UserTempMatch

func (uu *UserTempMatches) Build() []*models.UserTempMatch {
	var sUsers []*models.UserTempMatch

	for _, u := range *uu {
		swaggerUser := u.Build()
		sUsers = append(sUsers, &swaggerUser)
	}
	return sUsers
}
