package go_note

import "github.com/uptrace/bun"

type Note struct {
	bun.BaseModel `bun:"table:UserNotes"`

	NoteID    uint32 `json:"noteID" bun:"noteId,pk,autoincrement"`
	UserID    uint32 `json:"userID" bun:"userId,notnull"`
	NoteTitle string `json:"noteTitle" bun:"noteTitle"`
	NoteText  string `json:"noteText" bun:"noteText"`
}
