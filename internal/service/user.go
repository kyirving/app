package service

import (
	"app/internal/model"
	"app/internal/repository"
	"errors"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Login(user model.User) (model.User, error) {
	return s.Repo.Login(user)
}

func (s *UserService) Register(user model.User) (model.User, error) {
	exist, err := s.Repo.GetByUsername(user.Username)
	if err == nil && exist.ID != 0 {
		return model.User{}, errors.New("Username already exists")
	}

	return s.Repo.Register(user)
}
