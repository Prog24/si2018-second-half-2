package entities

import (
	"github.com/eure/si2018-second-half-2/models"
	"github.com/go-openapi/strfmt"
)

type UserImage struct {
	UserID    int64           `xorm:"user_id"`
	Path      string          `xorm:"path"`
	CreatedAt strfmt.DateTime `xorm:"created_at"`
	UpdatedAt strfmt.DateTime `xorm:"updated_at"`
}

func (u UserImage) Build() models.UserImage {
	return models.UserImage{
		UserID:    u.UserID,
		Path:      u.Path,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
