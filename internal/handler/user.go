package handler

import (
	"app/internal/dao"
	"app/internal/model"
	"app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{Svc: svc}
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dao.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	h.Svc.Login(user)
}
