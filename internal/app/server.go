package app

import (
	"app/config"
	"app/router"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	Config     *config.Config
	Db         *gorm.DB
	Redis      *redis.Client
	Web        *gin.Engine
	httpServer *http.Server
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	return &Server{
		Config: cfg,
		Db:     db,
		Web:    gin.Default(),
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%s", s.Config.App.Host, s.Config.App.Port)

	router.RegisterRouters(s.Web, s.Db, s.Config)

	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: s.Web,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
