package repositories

import (
	"github.com/eure/si2018-second-half-2/entities"
	"github.com/go-openapi/strfmt"
	"github.com/go-xorm/builder"
	"time"
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

// me_idで有効なレコードを探す
// 有効なレコードとは？？
func (r *UserTempMatchRepository) GetByUserID(userID int64) (*entities.UserTempMatch, error) {
	var ent = entities.UserTempMatch{}

	now := time.Now()
	startTime := strfmt.DateTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local))
	endTime := strfmt.DateTime(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 99, time.Local))

	s := r.GetSession()

	has, err := s.
		Where("user_id = ?", userID).
		Or("partner_id = ?", userID).
		And("created_at > ?", startTime).
		And("created_at < ?", endTime).
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

func (r *UserTempMatchRepository) GetLatest(userID int64, createdAt strfmt.DateTime) (*entities.UserTempMatch, error) {
	var ent = entities.UserTempMatch{}
	s := r.GetSession()
	has, err := s.Where("user_id = ?", userID).Or("partner_id = ?", userID).And("created_at = ?", createdAt).Get(&ent)
	if err != nil {
		return nil, err
	}
	if has {
		return &ent, nil
	}

	return nil, nil
}
