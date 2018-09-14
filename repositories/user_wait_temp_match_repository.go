package repositories

import (
	"fmt"
	"time"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/go-openapi/strfmt"
	"log"
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

func (r *UserWaitTempMatchRepository) Update(ent *entities.UserWaitTempMatch) error {
	now := strfmt.DateTime(time.Now())
	s := r.GetSession().Where("user_id = ?", ent.UserID).And("created_at = ?", ent.CreatedAt)
	ent.UpdatedAt = now
	defer func() { log.Println(s.LastSQL()) }()
	if _, err := s.Cols("user_id", "gender", "is_matched", "is_canceled", "created_at", "updated_at").Update(ent); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *UserWaitTempMatchRepository) IsMatchedToday(userID int64) (bool, error) {
	var ent entities.UserWaitTempMatch

	now := time.Now()
	startTime := strfmt.DateTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local))
	endTime := strfmt.DateTime(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 99, time.Local))

	s := r.GetSession()
	has, err := s.
		Where("user_id = ?", userID).
		And("created_at > ?", startTime).
		And("created_at < ?", endTime).
		And("is_matched = ?", true).
		Get(&ent)
	if err != nil {
		return false, err
	}
	if has {
		return true, nil
	}

	return false, nil
}

func (r *UserWaitTempMatchRepository) GetActive(user entities.User) (*entities.UserWaitTempMatch, error) {
	var ent entities.UserWaitTempMatch
	now := time.Now()
	limitTime := strfmt.DateTime(now.Add(time.Duration(-40) * time.Second))

	s := r.GetSession()
	has, err := s.
		Where("user_id = ?", user.ID).
		And("gender = ?", user.Gender).
		And("is_matched = ?", false).
		And("is_canceled = ?", false).
		And("created_at > ?", limitTime).
		Get(&ent)

	if err != nil {
		return nil, err
	}
	if has {
		return &ent, nil
	}
	return nil, nil
}

func (r *UserWaitTempMatchRepository) SearchPartner(user entities.User) (partnerID int64, err error) {
	var ent entities.UserWaitTempMatch
	oppositeGender := user.GetOppositeGender()

	now := time.Now()
	limitTime := strfmt.DateTime(now.Add(time.Duration(-40) * time.Second))
	fmt.Println(now)
	fmt.Println(limitTime)

	s := r.GetSession()
	// NOTE: Select が必要かどうか検証必要
	has, err := s.
		Select("user_wait_temp_match.*, residence_state").
		Join("INNER", "user", "user.id = user_wait_temp_match.user_id").
		Where("user_wait_temp_match.gender = ?", oppositeGender).
		And("residence_state = ?", user.ResidenceState).
		And("is_matched = ?", false).
		And("is_canceled = ?", false).
		And("user_wait_temp_match.created_at > ?", limitTime).
		Get(&ent)

	if err != nil {
		return 0, err
	}
	if has {
		return ent.UserID, nil
	}
	return 0, nil
}

func (r *UserWaitTempMatchRepository) GetLatestByUserID(userID int64) (*entities.UserWaitTempMatch, error) {
	var ent entities.UserWaitTempMatch

	s := r.GetSession()
	has, err := s.
		Where("user_id = ?", userID).
		Desc("created_at").
		Limit(1, 0).
		Get(&ent)

	if err != nil {
		return nil, err
	}
	if has {
		return &ent, nil
	}

	return nil, nil
}
