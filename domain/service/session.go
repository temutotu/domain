package service

import (
	"hello/domain/server/session"
	"net/http"
)

type Session struct {
}

func NewSessionService() (*Session, error) {
	service := &Session{}

	return service, nil
}

func (self *Session) CreateSession() (*http.Cookie, error) {
	sessionID, err := session.Create()
	if err != nil {
		return nil, err
	}

	cokkie := &http.Cookie{
		Name:  "session",
		Value: sessionID,
	}

	return cokkie, nil
}

func (self *Session) CheckSession(r *http.Request) bool {
	_, err := r.Cookie(session.CokkieName)
	if err == nil {
		return true
	} else {
		return false
	}
}
