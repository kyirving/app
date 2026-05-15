package service

import (
	"app/internal/model"
	"app/internal/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Login(username, password string) (model.User, error) {
	user, err := s.Repo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("invalid username or password")
		}
		return model.User{}, fmt.Errorf("login failed: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return model.User{}, errors.New("invalid username or password")
	}

	return user, nil
}

func (s *UserService) Register(user model.User) (model.User, error) {
	exist, err := s.Repo.GetByUsername(user.Username)
	if err == nil && exist.ID != 0 {
		return model.User{}, errors.New("username already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, fmt.Errorf("register failed: %w", err)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("register failed: %w", err)
	}
	user.Password = string(hashed)

	return s.Repo.Create(user)
}

func (s *UserService) GetByUserID(userID uint64) (model.User, error) {
	return s.Repo.GetByUserID(userID)
}
