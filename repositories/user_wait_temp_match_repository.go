package repositories

import (
	"github.com/eure/si2018-second-half-2/entities"
	"github.com/go-xorm/builder"
)

type UserWaitTempMatchRepository struct {
	RootRepository
}

func NewUserWaitMatchRepository(s *Session) UserWaitTempMatchRepository {
	return UserWaitTempMatchRepository{NewRootRepository(s)}
}

func (r *UserWaitTempMatchRepository) Create(ent entities.UserWaitTempMatch) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}
	return nil
}

func (r *UserWaitTempMatchRepository) Get(userID, partnerID int64) (*entities.UserWaitTempMatch, error) {
	var ent = entities.UserWaitTempMatch{}
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
