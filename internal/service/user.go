package service

import (
	"app/internal/model"
	"app/internal/repository"
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
