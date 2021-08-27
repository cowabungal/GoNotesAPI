package service

import (
	"GoNotes"
	"GoNotes/pkg/repository"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user *GoNotes.User) error {
	// тут можно генерить пассвордхэш
	return s.repo.Authorization.CreateUser(user)
}

func (s *AuthService) CreateSession(user *GoNotes.User) (*GoNotes.User, error) {
	user, err := s.repo.Authorization.CreateSession(user)

	return user, err
}

func (s *AuthService) CheckSession(id int) error {
	return s.repo.Authorization.CheckSession(id)
}
