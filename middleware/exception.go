package middleware

import (
	"app/pkg/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, resp.Output(resp.RESP_FAIL, nil, "Internal server error"))
		}
	}
}
