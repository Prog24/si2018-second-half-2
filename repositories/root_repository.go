package repositories

import (
	"errors"

	"github.com/go-xorm/xorm"
)

type RootRepository struct {
	s *Session
}

func NewRootRepository(s *Session) RootRepository {
	return RootRepository{s}
}

func (r *RootRepository) GetSession() *xorm.Session {
	if r.s.Session == nil {
		r.s.SetSession(engine.NewSession())
	}
	return r.s.Session
}

func TransactionBegin(s *Session) error {
	if s.Session == nil {
		s.SetSession(engine.NewSession())
	}
	return s.Session.Begin()
}

func TransactionRollBack(s *Session) error {
	if s.Session == nil {
		return errors.New("There is no transaction")
	}
	return s.Session.Rollback()
}

func TransactionCommit(s *Session) error {
	if s.Session == nil {
		return errors.New("There is no transaction")
	}
	return s.Session.Commit()
}
