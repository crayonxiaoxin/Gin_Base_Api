package controllers

import (
	"hello_gin_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Title       Login
// @Summary     登入
// @Description 登入
// @Param       username query string true "用户名"
// @Param       password query string true "密码"
// @Tags        登入
// @Success     200 {object} utils.Result
// @Failure     403 user     not exist
// @router      /login [post]
func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	result := models.Login(&models.User{Username: username, Password: password})
	ctx.JSON(http.StatusOK, result)
}
