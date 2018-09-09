package repositories

import "github.com/go-xorm/xorm"

type Session struct {
	Session *xorm.Session
}

func NewSession() *Session {
	return &Session{}
}

func (c *Session) SetSession(s *xorm.Session) {
	c.Session = s
}
