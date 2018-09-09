package token

import (
	"github.com/eure/si2018-second-half-2/entities"
	"github.com/eure/si2018-second-half-2/repositories"
)

func GetUserByToken(s *repositories.Session, token string) (*entities.User, error) {
	ent, err := repositories.NewUserTokenRepository(s).GetByToken(token)
	if err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, nil
	}

	user, err := repositories.NewUserRepository(s).GetByUserID(ent.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	return user, nil
}
