package handler

import (
	"app/internal/dao"
	"app/internal/model"
	"app/internal/service"
	"app/pkg/resp"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{Svc: svc}
}

// /user/login 登录
func (u *UserHandler) Login(c *gin.Context) {
	var req dao.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	u.Svc.Login(user)
}

// /user/register 注册
func (u *UserHandler) Register(c *gin.Context) {
	var req dao.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp.Output(resp.RESP_BAD_REQUEST, nil, "Register params failed"))
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	user, err := u.Svc.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Output(resp.RESP_FAIL, nil, err.Error()))
		return
	}
	fmt.Println("userid", user.ID)
	c.JSON(http.StatusOK, resp.Output(resp.RESP_SUCC, user, "Register success"))

}
