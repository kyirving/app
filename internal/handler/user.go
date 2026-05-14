package handler

import (
	"app/internal/dao"
	"app/internal/model"
	"app/internal/service"
	"app/pkg/jwt"
	"app/pkg/resp"
	"fmt"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Svc          *service.UserService
	Snowflake    *snowflake.Node
	AccessSecret string
	RefreshSecret string
}

func NewUserHandler(svc *service.UserService, snowflakeNode *snowflake.Node, accessSecret, refreshSecret string) *UserHandler {
	return &UserHandler{
		Svc:           svc,
		Snowflake:     snowflakeNode,
		AccessSecret:  accessSecret,
		RefreshSecret: refreshSecret,
	}
}

// POST /user/login 登录
func (u *UserHandler) Login(c *gin.Context) {
	var req dao.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp.Output(resp.RESP_BAD_REQUEST, nil, "Login params failed"))
		return
	}

	user, err := u.Svc.Login(model.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, resp.Output(resp.RESP_UNAUTHORIZED, nil, "Invalid username or password"))
		return
	}

	token, err := jwt.GenerateToken(user.UserID, "user", u.AccessSecret, u.RefreshSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Output(resp.RESP_FAIL, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.Output(resp.RESP_SUCC, dao.LoginResponse{
		Token: token,
		User: dao.UserInfo{
			UserID:   user.UserID,
			Username: user.Username,
			Mobile:   user.Mobile,
			Nikename: user.Nikename,
		},
	}, "Login success"))
}

// POST /user/register 注册
func (u *UserHandler) Register(c *gin.Context) {
	var req dao.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp.Output(resp.RESP_BAD_REQUEST, nil, "Register params failed"))
		return
	}

	user := model.User{
		UserID:   uint64(u.Snowflake.Generate().Int64()),
		Username: req.Username,
		Password: req.Password,
	}
	_, err := u.Svc.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Output(resp.RESP_FAIL, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.Output(resp.RESP_SUCC, nil, "Register success"))
}

// GET /user/info 获取用户信息（需鉴权）
func (u *UserHandler) GetInfo(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, resp.Output(resp.RESP_UNAUTHORIZED, nil, "Unauthorized"))
		return
	}
	userClaims := claims.(*jwt.UserClaims)

	user, err := u.Svc.GetByUserID(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, resp.Output(resp.RESP_NOT_FOUND, nil, fmt.Sprintf("User not found: %v", err)))
		return
	}

	c.JSON(http.StatusOK, resp.Output(resp.RESP_SUCC, dao.UserInfo{
		UserID:   user.UserID,
		Username: user.Username,
		Mobile:   user.Mobile,
		Nikename: user.Nikename,
	}, "success"))
}
