package service

import (
	"GoNotes"
	"GoNotes/pkg/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserInfo(id int) (*GoNotes.UserInfo, error) {
	return s.repo.User.GetUserInfo(id)
}
