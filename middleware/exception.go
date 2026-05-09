package middleware

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			// 判断错误类型
			fmt.Println(" err typeof", reflect.TypeOf(err))
		}
	}
}
