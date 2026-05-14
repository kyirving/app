package dao

import "app/pkg/jwt"

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
