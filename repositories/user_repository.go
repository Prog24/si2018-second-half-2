package repositories

import (
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/eure/si2018-second-half-2/entities"
)

type UserRepository struct {
	RootRepository
}

func NewUserRepository(s *Session) *UserRepository {
	return &UserRepository{NewRootRepository(s)}
}

func (r *UserRepository) Create(ent entities.User) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(ent *entities.User) error {
	now := strfmt.DateTime(time.Now())
	s := r.GetSession().Where("id = ?", ent.ID)
	ent.UpdatedAt = now
	if _, err := s.Update(ent); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByUserID(userID int64) (*entities.User, error) {
	var ent = entities.User{ID: userID}

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

// limit / offset / 検索対象の性別 でユーザーを取得
// idsには取得対象に含めないUserIDを入れる (いいね/マッチ/ブロック済みなど)
func (r *UserRepository) FindWithCondition(limit, offset int, gender string, ids []int64) ([]entities.User, error) {
	var users []entities.User

	s := r.GetSession()

	s.Where("gender = ?", gender)
	if len(ids) > 0 {
		s.NotIn("id", ids)
	}
	s.Limit(limit, offset)
	s.Desc("id")

	err := s.Find(&users)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *UserRepository) FindByIDs(ids []int64) ([]entities.User, error) {
	var users []entities.User

	err := engine.In("id", ids).Find(&users)
	if err != nil {
		return users, err
	}

	return users, nil
}
