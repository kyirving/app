package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwtV5 "github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID uint
	Role   string
	jwtV5.RegisteredClaims
}

type Token struct {
	UserID           uint      `json:"user_id"`
	AccessToken      string    `json:"access_token"`
	ExpiresAt        time.Time `json:"expires_at"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
}

func GenerateToken(userID uint, role, secretKey, refreshSecretKey string) (Token, error) {

	// token 过期时间 24 小时
	claims := UserClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "wuh",       // 签发者
			Subject:   "user-auth", // 主题
		},
	}

	claims2 := UserClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "wuh",       // 签发者
			Subject:   "user-auth", // 主题
		},
	}
	Token := Token{
		UserID: userID,
	}
	// access token 过期时间 24 小时
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return Token, err
	}
	Token.AccessToken = accessToken
	Token.ExpiresAt = claims.ExpiresAt.Time

	// refresh token 过期时间 7 天
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	refreshToken, err := token2.SignedString(refreshSecretKey)
	if err != nil {
		return Token, err
	}
	Token.RefreshToken = refreshToken
	Token.RefreshExpiresAt = claims2.ExpiresAt.Time
	return Token, nil
}
