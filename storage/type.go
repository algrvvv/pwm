package storage

import "time"

type Note struct {
	ID          int
	Name        string
	Value       string
	UsePassword bool
	CreatedAt   time.Time
}

type Storage interface {
	GetNotes() (notes []Note, err error)
	GetNoteByName(name string) (note Note, err error)
	SaveNote(note Note) (err error)
	DeleteNoteByName(name string) (err error)
	Migrate() (err error)
}
