package service

import (
	"GoNotes"
	"GoNotes/pkg/repository"
)

type Service struct {
	Authorization
	User
	Note
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repo), User: NewUserService(repo), Note: NewNoteService(repo)}
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
	Update(userId int, note *GoNotes.Note) (int, error)
}
