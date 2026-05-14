package middleware

import (
	"app/pkg/jwt"
	"app/pkg/resp"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, resp.Output(resp.RESP_UNAUTHORIZED, nil, "Missing authorization header"))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, resp.Output(resp.RESP_UNAUTHORIZED, nil, "Invalid authorization format"))
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(parts[1], secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, resp.Output(resp.RESP_UNAUTHORIZED, nil, "Invalid or expired token"))
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
