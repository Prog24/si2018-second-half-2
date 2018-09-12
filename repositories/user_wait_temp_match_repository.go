package repositories

import (
	"github.com/eure/si2018-second-half-2/entities"
	"github.com/go-xorm/builder"
)

type UserWaitTempMatchRepository struct {
	RootRepository
}

func NewUserWaitTempMatchRepository(s *Session) UserWaitTempMatchRepository {
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

func (r *UserWaitTempMatchRepository) IsIActive(user entities.User) (bool, error) {
	var ent entities.UserWaitTempMatch

	s := r.GetSession()
	has, err := s.Where("user_id = ?", user.ID).And("gender = ?", user.Gender).And("is_matched = ?", false).And("is_canceled = ?", false).Get(&ent)
	if err != nil {
		return false, err
	}
	if has {
		return true, nil
	}
	return false, nil
}

func (r *UserWaitTempMatchRepository) SearchPartner(user entities.User) (partnerID int64, err error) {
	var ent entities.UserWaitTempMatch
	oppositeGender := user.GetOppositeGender()

	s := r.GetSession()
	// NOTE: Select が必要かどうか検証必要
	has, err := s.
		Select("user_wait_temp_match.*, residence_state").
		Join("INNER", "user", "user.id = user_wait_temp_match.user_id").
		Where("gender = ?", oppositeGender).
		And("residence_state = ?", user.ResidenceState).
		And("is_matched = ?", false).
		And("is_canceled = ?", false).
		Get(&ent)

	if err != nil {
		return 0, err
	}
	if has {
		return ent.UserID, nil
	}
	return 0, nil
}
