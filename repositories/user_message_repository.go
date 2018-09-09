package repositories

import (
	"log"

	"github.com/eure/si2018-second-half-2/entities"
	"github.com/go-openapi/strfmt"
	"github.com/go-xorm/builder"
)

type UserMessageRepository struct {
	RootRepository
}

func NewUserMessageRepository(s *Session) UserMessageRepository {
	return UserMessageRepository{NewRootRepository(s)}
}

func (r *UserMessageRepository) Create(ent entities.UserMessage) error {
	s := r.GetSession()
	if _, err := s.Insert(&ent); err != nil {
		return err
	}

	return nil
}

// userとpartnerがやりとりしたメッセージをlimit/latest/oldestで取得する.
func (r *UserMessageRepository) GetMessages(userID, partnerID int64, limit int, latest, oldest *strfmt.DateTime) ([]entities.UserMessage, error) {
	var messages []entities.UserMessage
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
