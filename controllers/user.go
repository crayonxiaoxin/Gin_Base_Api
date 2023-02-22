package controllers

import (
	"fmt"
	"hello_gin_api/models"
	"hello_gin_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Title			GetUsers
//	@Summary		获取所有用户
//	@Description	获取所有用户
//	@Param			page	query	int	false	"页码"
//	@Param			size	query	int	false	"每页数量"
//	@Tags			用户相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/user [get]
func GetUsers(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageInt := utils.Str2Int(page)
	sizeInt := utils.Str2Int(size)
	users, count := models.GetUsers(int(pageInt), int(sizeInt))
	data := make(map[string]interface{})
	data["list"] = users
	data["count"] = count
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: data}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			GetUser
//	@Summary		通过id获取用户
//	@Description	通过id获取用户
//	@Param			id	path	int	true	"用户id"
//	@Tags			用户相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/user/{id} [get]
func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := utils.Str2Int(id, 0)
	var result = utils.Result{}
	if uid > 0 {
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

//	@Title			CreateUser
//	@Summary		添加用户
//	@Description	添加用户
//	@Param			user_login	query	string	true	"用户名"
//	@Param			user_pass	query	string	true	"密码"
//	@Tags			用户相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/user [post]
func AddUser(ctx *gin.Context) {
	user_login := ctx.Query("user_login")
	user_pass := ctx.Query("user_pass")
	result := models.AddUser(&models.User{UserLogin: user_login, UserPass: user_pass})
	ctx.JSON(http.StatusOK, result)
}

//	@Title			UpdateUser
//	@Summary		更新用户
//	@Description	更新用户
//	@Param			id			path	int		true	"用户id"
//	@Param			user_login	query	string	true	"用户名"
//	@Param			user_pass	query	string	false	"密码"
//	@Tags			用户相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/user/{id} [put]
func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := utils.Str2Int(id, 0)
	user_login := ctx.Query("user_login")
	user_pass := ctx.Query("user_pass")
	user := models.User{UserLogin: user_login, UserPass: user_pass}
	user.ID = uint(uid)
	result := models.UpdateUser(&user)
	ctx.JSON(http.StatusOK, result)
}

//	@Title			DeleteUser
//	@Summary		删除用户
//	@Description	删除用户
//	@Param			id	path	int	true	"用户id"
//	@Tags			用户相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/user/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("id: %v\n", id)
	uid := utils.Str2Int(id, 0)
	result := models.DeleteUser(int(uid))
	ctx.JSON(http.StatusOK, result)
}

//	@Title			Get UserMetas
//	@Summary		通过uid获取元数据
//	@Description	通过uid获取元数据
//	@Param			id			path	int		true	"用户id"
//	@Param			meta_key	query	string	false	"Key，如果填写，只返回对应值"
//	@Tags			用户相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/user/{id}/meta [get]
func GetUserMetas(ctx *gin.Context) {
	id := ctx.Param("id")
	meta_key := ctx.Query("meta_key")
	uid := utils.Str2Int(id, 0)
	var result = utils.Result{}
	if uid > 0 {
		if len(meta_key) > 0 { // 如果有 meta_key，则获取单个值
			m := make(map[string]string)
			m[meta_key] = models.GetUserMetaValue(&models.UserMeta{Uid: uint(uid), MetaKey: meta_key})
			result.Data = m
		} else { // 否则，获取所有相关值
			m := models.GetUserMetas(uid)
			result.Data = m
		}
		result.ResultCode = utils.SUCCESS
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			Update UserMeta
//	@Summary		新增或更新用户元数据
//	@Description	新增或更新用户元数据
//	@Param			id			path	int		true	"用户id"
//	@Param			meta_key	query	string	true	"Key"
//	@Param			meta_value	query	string	false	"Value"
//	@Tags			用户相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/user/{id}/meta [post]
func UpdateUserMeta(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := utils.Str2Int(id, 0)
	meta_key := ctx.Query("meta_key")
	meta_value := ctx.Query("meta_value")
	meta := models.UserMeta{Uid: uint(uid), MetaKey: meta_key, MetaValue: meta_value}
	result := models.UpdateUserMeta(&meta)
	ctx.JSON(http.StatusOK, result)
}

//	@Title			DeleteUserMeta
//	@Summary		删除用户元数据
//	@Description	删除用户元数据
//	@Param			id			path	int		true	"用户id"
//	@Param			meta_key	query	string	true	"Key"
//	@Tags			用户相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/user/{id}/meta [delete]
func DeleteUserMeta(ctx *gin.Context) {
	id := ctx.Param("id")
	meta_key := ctx.Query("meta_key")
	uid := utils.Str2Int(id, 0)
	result := models.DeleteUserMeta(uint(uid), meta_key)
	ctx.JSON(http.StatusOK, result)
}
