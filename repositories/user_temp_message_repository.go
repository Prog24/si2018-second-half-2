package repositories

import (
	"log"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/go-openapi/strfmt"
	"github.com/go-xorm/builder"
)

type UserTempMessageRepository struct {
	RootRepository
}

func NewUserTempMessageRepository(s *Session) UserTempMessageRepository {
	return UserTempMessageRepository{NewRootRepository(s)}
}

func (r *UserTempMessageRepository) Create(ent entities.UserTempMessage) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

// userとpartnerがやりとりしたメッセージをlimit/latest/oldestで取得する.
func (r *UserTempMessageRepository) GetMessages(userID, partnerID int64, limit int, latest, oldest *strfmt.DateTime) ([]entities.UserTempMessage, error) {
	var messages []entities.UserTempMessage
	var ids = []int64{userID, partnerID}

	s := r.GetSession()
	defer func() { log.Println(s.LastSQL()) }()
	s.Where(builder.In("user_id", ids))
	s.And(builder.In("partner_id", ids))
	if latest != nil {
		s.And("created_at < ?", latest)
	}
	if oldest != nil {
		s.And("created_at > ?", oldest)
	}
	s.Desc("created_at")
	s.Limit(limit)
	err := s.Find(&messages)
	if err != nil {
		return messages, err
	}

	return messages, nil
}
