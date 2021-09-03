package service

import (
	"GoNotes"
	"GoNotes/pkg/repository"
	"crypto/sha1"
	"fmt"
	"os"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user GoNotes.User) error {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.Authorization.CreateUser(user)
}

func (s *AuthService) CreateSession(user *GoNotes.User) (*GoNotes.User, error) {
	user.Password = generatePasswordHash(user.Password)
	user, err := s.repo.Authorization.CreateSession(user)

	return user, err
}

func (s *AuthService) CheckSession(id int) error {
	return s.repo.Authorization.CheckSession(id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
