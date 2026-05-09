package repository

import (
	"app/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u *UserRepository) Login(user model.User) (model.User, error) {
	var result model.User
	if err := u.Db.Where("username = ? AND password = ?", user.Username, user.Password).First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
