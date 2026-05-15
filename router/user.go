package router

import (
	"app/config"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"app/middleware"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRouter(r *gin.Engine, db *gorm.DB, jwtCfg config.JwtConfig) {
	snowflakeNode, err := snowflake.NewNode(jwtCfg.SnowflakeNodeID)
	if err != nil {
		panic(fmt.Sprintf("failed to init snowflake: %v", err))
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, snowflakeNode, jwtCfg.AccessSecret, jwtCfg.RefreshSecret)

	userRouter := r.Group("/user")
	{
		userRouter.POST("/login", middleware.RateLimitMiddleware(5, time.Second*10), userHandler.Login)
		userRouter.POST("/register", middleware.RateLimitMiddleware(3, time.Second*10), userHandler.Register)

		auth := userRouter.Group("")
		auth.Use(middleware.AuthMiddleware(jwtCfg.AccessSecret))
		{
			auth.GET("/info", userHandler.GetInfo)
		}
	}
}
