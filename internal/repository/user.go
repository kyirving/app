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

func (u *UserRepository) GetByUsername(username string) (model.User, error) {
	var result model.User
	if err := u.Db.Where("username = ?", username).First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (u *UserRepository) Register(user model.User) (model.User, error) {
	if err := u.Db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) GetByUserID(userID uint64) (model.User, error) {
	var result model.User
	if err := u.Db.Where("user_id = ?", userID).First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
