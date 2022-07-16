package handlers

import (
	"encoding/json"
	"fmt"
	"go-notes-webapp/main-module/dbmanager"
	"go-notes-webapp/main-module/encryption"
	"go-notes-webapp/main-module/filemanager"
	"go-notes-webapp/main-module/go_user"
	"strings"
	"testing"
)

func TestHandleLoginRequest(t *testing.T) {
	url, err := filemanager.ReadFile("/etc/server/c/r")
	if err != nil {
		t.Error(err)
	}
	var dbm dbmanager.DatabaseManager
	err = dbm.Connect(strings.Trim(string(url), "\n"))
	if err != nil {
		t.Error(err)
	}

	p, err := encryption.EncryptSHA([]byte("pass1"))
	if err != nil {
		t.Error(err)
	}
	//
	u := go_user.User{ID: 1, Username: "user1", Password: p}
	//err = dbm.Insert(&u)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//notes := []go_note.Note{
	//	{UserID: 1, NoteText: "test1"},
	//	{UserID: 1, NoteText: "rrrr"},
	//	{UserID: 1, NoteText: "wqdfqwf"},
	//}
	//
	//err = dbm.Insert(&notes)
	//if err != nil {
	//	t.Error(err)
	//}

	//var n []go_note.Note
	err = dbm.Select(&u.Notes, "userId=?", u.ID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
	j, err := json.Marshal(u)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(j))
	fmt.Println(strings.ReplaceAll(string(j), `"`, "'"))
}
