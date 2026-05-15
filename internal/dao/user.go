package dao

import "app/pkg/jwt"

type Login struct {
	Username string `json:"username" binding:"required,min=2,max=32"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type Register struct {
	Username string `json:"username" binding:"required,min=2,max=32"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type LoginResponse struct {
	Token jwt.Token `json:"token"`
	User  UserInfo  `json:"user"`
}

type UserInfo struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile,omitempty"`
	Nikename string `json:"nikename,omitempty"`
}
