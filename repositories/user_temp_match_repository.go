package repositories

import (
	"github.com/eure/si2018-second-half-2/entities"
	"github.com/go-xorm/builder"
)

type UserTempMatchRepository struct {
	RootRepository
}

func NewUserTempMatchRepository(s *Session) UserTempMatchRepository {
	return UserTempMatchRepository{NewRootRepository(s)}
}

func (r *UserTempMatchRepository) Create(ent entities.UserTempMatch) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *UserTempMatchRepository) Get(userID, partnerID int64) (*entities.UserTempMatch, error) {
	var ent = entities.UserTempMatch{}
	var ids = []int64{userID, partnerID}
	s := r.GetSession()
	has, err := s.Where(builder.In("user_id", ids).And(builder.In("partner_id", ids))).Get(&ent)
	if err != nil {
		return nil, err
	}
	if has {
		return &ent, nil
	}
	return nil, nil
}
