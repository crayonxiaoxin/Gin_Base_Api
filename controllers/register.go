package controllers

import (
	"hello_gin_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Title       Register
// @Description 注册
// @Param       username query string true "用户名"
// @Param       password query string true "密码"
// @Tags        register
// @Success     200 {object} utils.Result
// @Failure     403 user     not exist
// @router      /register [post]
func Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	result := models.Register(&models.User{Username: username, Password: password})
	ctx.JSON(http.StatusOK, result)
}
