package go_user

import (
	"fmt"
	"github.com/uptrace/bun"
	"go-notes-webapp/main-module/go_note"
)

type User struct {
	bun.BaseModel `bun:"table:Users"`

	ID       uint32         `json:"id" bun:"userId,pk,autoincrement"`
	Username string         `json:"username" bun:"username,unique,notnull"`
	Password []byte         `json:"-" bun:"password,notnull"`
	Notes    []go_note.Note `json:"notes" bun:"-"`
}

func (u User) String() string {
	return fmt.Sprintf("go_notes.User struct data:\nuserId: %v\nusername: %v\npassword: %v\nuser notes: %v\n", u.ID, u.Username, u.Password, u.Notes)
}

type UserStringPass struct {
	//bun.BaseModel `bun:"table:Users"`

	ID       uint32         `json:"id" bun:"userId,pk,autoincrement"`
	Username string         `json:"username" bun:"username,unique,notnull"`
	Password string         `json:"password" bun:"password,notnull"`
	Notes    []go_note.Note `json:"notes" bun:"-"`
}
