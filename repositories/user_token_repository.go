package repositories

import "github.com/eure/si2018-second-half-2/entities"

type UserTokenRepository struct {
	RootRepository
}

func NewUserTokenRepository(s *Session) *UserTokenRepository {
	return &UserTokenRepository{NewRootRepository(s)}
}

func (r *UserTokenRepository) Create(ent entities.UserToken) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *UserTokenRepository) Update(ent entities.UserToken, cols []string) error {
	s := r.GetSession()
	s.MustCols(cols...)
	if _, err := s.Update(ent); err != nil {
		return err
	}
	return nil
}

func (r *UserTokenRepository) GetByUserID(userID int64) (*entities.UserToken, error) {
	var ent = entities.UserToken{UserID: userID}

	s := r.GetSession()
	has, err := s.Get(&ent)
	if err != nil {
		return nil, err
	}

	if has {
		return &ent, nil
	}

	return nil, nil
}

func (r *UserTokenRepository) GetByToken(token string) (*entities.UserToken, error) {
	var ent = entities.UserToken{Token: token}

	s := r.GetSession()
	has, err := s.Get(&ent)
	if err != nil {
		return nil, err
	}

	if has {
		return &ent, nil
	}

	return nil, nil
}
