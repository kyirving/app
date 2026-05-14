package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID uint64
	Role   string
	jwt.RegisteredClaims
}

type Token struct {
	UserID           uint64    `json:"user_id"`
	AccessToken      string    `json:"access_token"`
	ExpiresAt        time.Time `json:"expires_at"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
}

func GenerateToken(userID uint64, role, secretKey, refreshSecretKey string) (Token, error) {

	// access token 过期时间 24 小时
	claims := UserClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "wuh",
			Subject:   "user-auth",
		},
	}

	// refresh token 过期时间 7 天
	claims2 := UserClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "wuh",
			Subject:   "user-auth",
		},
	}
	Token := Token{
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return Token, err
	}
	Token.AccessToken = accessToken
	Token.ExpiresAt = claims.ExpiresAt.Time

	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	refreshToken, err := token2.SignedString([]byte(refreshSecretKey))
	if err != nil {
		return Token, err
	}
	Token.RefreshToken = refreshToken
	Token.RefreshExpiresAt = claims2.ExpiresAt.Time
	return Token, nil
}

func ParseToken(tokenString, secretKey string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
