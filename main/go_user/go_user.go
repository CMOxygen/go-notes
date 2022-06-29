package go_user

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:Users"`

	ID       uint32 `json:"id" bun:"userId,pk,autoincrement"`
	Username string `json:"username" bun:"username,unique,notnull"`
	Password string `json:"password" bun:"password,notnull"`
}
