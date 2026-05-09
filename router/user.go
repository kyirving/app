package router

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRouter(r *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userRouter := r.Group("/user")
	{
		userRouter.POST("/login", userHandler.Login)
	}
}
