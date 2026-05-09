package app

import (
	"app/config"
	"app/router"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	Config *config.Config
	Db     *gorm.DB
	Redis  *redis.Client
	Web    *gin.Engine
}

func NewServer(config *config.Config, db *gorm.DB) *Server {
	return &Server{
		Config: config,
		Db:     db,
		Web:    gin.Default(),
	}
}

func (s *Server) Start() {
	addr := fmt.Sprintf("%s:%s", s.Config.App.Host, s.Config.App.Port)
	// 注册路由

	router.RegisterRouters(s.Web, s.Db)

	s.Web.Run(addr)
}
