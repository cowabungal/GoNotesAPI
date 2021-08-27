package repository

import (
	"GoNotes"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	User
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthRepository(db), User: NewUserRepository(db), Note: NewNoteRepository(db)}
}

type Authorization interface {
	CreateUser(user *GoNotes.User) error
	CreateSession(user *GoNotes.User) (*GoNotes.User, error)
	CheckSession(id int) error
}

type User interface {
	GetUserInfo(id int) (*GoNotes.UserInfo, error)
}

type Note interface {
	GetAll(userId int) ([]*GoNotes.Note, error)
	Add(id int, note *GoNotes.Note) (int, error)
	Delete(id, userId int) error
}
