package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouters(r *gin.Engine, db *gorm.DB) {
	//用户路由
	RegisterUserRouter(r, db)
}
