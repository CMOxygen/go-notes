package sessions

import (
	"fmt"
	"go-notes-webapp/main-module/go_user"
	"testing"
)

func TestGenerateSessionID(t *testing.T) {
	t.Log("testing generate sessions")
	{
		u := go_user.User{ID: 1}
		var sm SessionManager
		sm.CreateSession(u)
		fmt.Println(sm)
		fmt.Println(sm.SessionExists(sm.Sessions[0].SessionID))

		u2 := go_user.User{ID: 2}
		sm.CreateSession(u2)
		fmt.Println(sm)

		err := sm.StopSession(sm.Sessions[0].SessionID)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(sm)
	}
}
