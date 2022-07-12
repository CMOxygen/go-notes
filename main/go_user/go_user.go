package go_user

import (
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

type UserStringPass struct {
	//bun.BaseModel `bun:"table:Users"`

	ID       uint32         `json:"id" bun:"userId,pk,autoincrement"`
	Username string         `json:"username" bun:"username,unique,notnull"`
	Password string         `json:"password" bun:"password,notnull"`
	Notes    []go_note.Note `json:"notes" bun:"-"`
}
