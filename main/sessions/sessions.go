package sessions

import (
	"encoding/hex"
	"errors"
	"go-notes-webapp/main-module/go_user"
	"math/rand"
	"strings"
	"time"
)

type Session struct {
	SessionID string
	UserID    uint32
}

func (s *Session) GenerateSessionID() {

	alphabet := []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_=-+/*!?")
	var randomBytes []byte

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 32; i++ {
		randomBytes = append(randomBytes, alphabet[rand.Intn(len(alphabet))])
	}
	s.SessionID = hex.EncodeToString(randomBytes)
}

type SessionManager struct {
	Sessions []Session
}

func (m *SessionManager) CreateSession(u go_user.User) {
	var s Session
	s.GenerateSessionID()
	s.UserID = u.ID
	m.Sessions = append(m.Sessions, s)
}

func (m *SessionManager) StopSession(sessionId string) error {

	for i, s := range m.Sessions {
		if strings.Compare(s.SessionID, sessionId) == 0 {
			m.Sessions = append(m.Sessions[:i], m.Sessions[i+1:]...)
			return nil
		}
	}
	return errors.New("stop session: no session to stop")
}

func (m *SessionManager) SessionExists(sessionId string) bool {
	for _, s := range m.Sessions {
		if strings.Compare(s.SessionID, sessionId) == 0 {
			return true
		}
	}
	return false
}

//func (m *SessionManager) GetSession() Session {
//
//}
