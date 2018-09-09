package repositories

import (
	"time"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/go-openapi/strfmt"
)

type UserImageRepository struct {
	RootRepository
}

func NewUserImageRepository(s *Session) UserImageRepository {
	return UserImageRepository{NewRootRepository(s)}
}

func (r *UserImageRepository) Create(ent entities.UserImage) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

func (r *UserImageRepository) Update(ent entities.UserImage) error {
	now := strfmt.DateTime(time.Now())

	s := r.GetSession().Where("user_id = ?", ent.UserID)
	ent.UpdatedAt = now

	if _, err := s.Update(ent); err != nil {
		return err
	}
	return nil
}

func (r *UserImageRepository) GetByUserID(userID int64) (*entities.UserImage, error) {
	var ent = entities.UserImage{UserID: userID}

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

func (r *UserImageRepository) GetByUserIDs(userIDs []int64) ([]entities.UserImage, error) {
	var userImages []entities.UserImage

	s := r.GetSession()
	err := s.In("user_id", userIDs).Find(&userImages)
	if err != nil {
		return userImages, err
	}

	return userImages, nil
}
