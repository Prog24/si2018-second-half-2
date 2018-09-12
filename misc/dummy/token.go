package main

import (
	"fmt"
	"time"

	"github.com/eure/si2018-server-side/entities"
	"github.com/eure/si2018-server-side/repositories"
	"github.com/go-openapi/strfmt"
)

func dummyToken() {
	r := repositories.NewUserTokenRepository()

	for i := firstUserID; i <= lastUserID; i++ {
		now := strfmt.DateTime(time.Now())
		token := entities.UserToken{
			UserID:    int64(i),
			Token:     fmt.Sprintf("USERTOKEN%v", i),
			CreatedAt: now,
			UpdatedAt: now,
		}
		r.Create(token)
	}
}
