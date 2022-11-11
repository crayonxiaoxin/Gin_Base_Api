package controllers

import (
	"fmt"
	"hello_gin_api/models"
	"hello_gin_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Title       GetAll
// @Description 获取所有用户
// @Param       token header string true "登入后返回的token"
// @Tags        user
// @Success     200 {object} utils.Result
// @router      /user [get]
func GetAllUsers(ctx *gin.Context) {
	users := models.GetAllUsers()
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: users}
	ctx.JSON(http.StatusOK, result)
}

// @Title       Get
// @Description 通过id获取用户
// @Param       token header string true "登入后返回的token"
// @Param       id    path   int    true "The key for staticblock"
// @Tags        user
// @Success     200 {object} utils.Result
// @router      /user/{id} [get]
func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 0, 0)
	var result = utils.Result{}
	if err == nil && uid != 0 {
		user := models.GetUser(int(uid))
		if user.Valid() {
			result.ResultCode = utils.SUCCESS
			result.Data = user
		} else {
			result.ResultCode = utils.ERR_USER_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	ctx.JSON(http.StatusOK, result)
}

// @Title       CreateUser
// @Description 添加用户
// @Param       token    header string true "登入后返回的token"
// @Param       username query  string true "用户名"
// @Param       password query  string true "密码"
// @Tags        user
// @Success     200 {object} utils.Result
// @router      /user [post]
func AddUser(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	result := models.AddUser(&models.User{Username: username, Password: password})
	ctx.JSON(http.StatusOK, result)
}

// @Title       Update
// @Description 更新用户
// @Param       token    header string true  "登入后返回的token"
// @Param       id       path   int    true  "The uid you want to update"
// @Param       username query  string true  "用户名"
// @Param       password query  string false "密码"
// @Tags        user
// @Success     200 {object} utils.Result
// @router      /user/{id} [put]
func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		uid = 0
	}
	username := ctx.Query("username")
	password := ctx.Query("password")
	user := models.User{Username: username, Password: password}
	user.ID = uint(uid)
	result := models.UpdateUser(&user)
	ctx.JSON(http.StatusOK, result)
}

// @Title       Delete
// @Description 删除用户
// @Param       token header string true "登入后返回的token"
// @Param       id    path   int    true "The uid you want to delete"
// @Tags        user
// @Success     200 {object} utils.Result
// @router      /user/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("id: %v\n", id)
	uid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		uid = 0
	}
	result := models.DeleteUser(int(uid))
	ctx.JSON(http.StatusOK, result)
}
