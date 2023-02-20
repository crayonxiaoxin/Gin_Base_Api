package controllers

import (
	"hello_gin_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Title       Login
// @Summary     登入
// @Description 登入
// @Param       user_login query string true "用户名"
// @Param       user_pass query string true "密码"
// @Tags        登入
// @Success     200 {object} utils.Result
// @Failure     403 user     not exist
// @router      /login [post]
func Login(ctx *gin.Context) {
	user_login := ctx.Query("user_login")
	user_pass := ctx.Query("user_pass")
	result := models.Login(&models.User{UserLogin: user_login, UserPass: user_pass})
	ctx.JSON(http.StatusOK, result)
}
