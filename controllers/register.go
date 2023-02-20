package controllers

import (
	"hello_gin_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Title       Register
// @Summary     注册
// @Description 注册
// @Param       user_login query string true "用户名"
// @Param       user_pass query string true "密码"
// @Tags        注册
// @Success     200 {object} utils.Result
// @Failure     403 user     not exist
// @router      /register [post]
func Register(ctx *gin.Context) {
	user_login := ctx.Query("user_login")
	user_pass := ctx.Query("user_pass")
	result := models.Register(&models.User{UserLogin: user_login, UserPass: user_pass})
	ctx.JSON(http.StatusOK, result)
}
