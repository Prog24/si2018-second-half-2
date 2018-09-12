package entities

import (
	"github.com/eure/si2018-second-half-2/models"
	"github.com/go-openapi/strfmt"
)

type UserTempMessage struct {
	UserID    int64           `xorm:"user_id"`
	PartnerID int64           `xorm:"partner_id"`
	Message   string          `xorm:"message"`
	CreatedAt strfmt.DateTime `xorm:"created_at"`
	UpdatedAt strfmt.DateTime `xorm:"updated_at"`
}

func (u UserTempMessage) Build() models.UserTempMessage {
	return models.UserTempMessage{
		UserID:    u.UserID,
		PartnerID: u.PartnerID,
		Message:   u.Message,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type UserTempMessages []UserTempMessage

func (msgs *UserTempMessages) Build() []*models.UserTempMessage {
	var sMsgs []*models.UserTempMessage

	for _, m := range *msgs {
		sMsg := m.Build()
		sMsgs = append(sMsgs, &sMsg)
	}
	return sMsgs
}
