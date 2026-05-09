package router

import (
	"app/middleware"
	"app/pkg/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouters(r *gin.Engine, db *gorm.DB) {
	r.Use(middleware.ExceptionMiddleware())
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, resp.Output(resp.RESP_METHOD_NOT_ALLOWED, nil, "Method Not Allowed"))
		c.Abort()
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, resp.Output(resp.RESP_NOT_FOUND, nil, "Not Found"))
		c.Abort()
	})

	//用户路由
	RegisterUserRouter(r, db)
}
