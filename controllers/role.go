package controllers

import (
	"hello_gin_api/models"
	"hello_gin_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Title			GetRoles
//	@Summary		获取角色
//	@Description	获取角色
//	@Param			page	query	int	false	"页码"
//	@Param			size	query	int	false	"每页数量"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/role [get]
func GetRoles(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageInt := utils.Str2Int(page)
	sizeInt := utils.Str2Int(size)
	users, count := models.GetRoles(int(pageInt), int(sizeInt))
	data := make(map[string]interface{})
	data["list"] = users
	data["count"] = count
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: data}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			GetRole
//	@Summary		通过id获取角色
//	@Description	通过id获取角色
//	@Param			id	path	int	true	"角色id"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/role/{id} [get]
func GetRole(ctx *gin.Context) {
	id := ctx.Param("id")
	post_id := utils.Str2Int(id, 0)
	var result = utils.Result{}
	if post_id > 0 {
		post := models.GetRole(int(post_id))
		if post.Valid() {
			result.ResultCode = utils.SUCCESS
			result.Data = post
		} else {
			result.ResultCode = utils.ERR_POST_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			AddRole
//	@Summary		添加角色
//	@Description	添加角色
//	@Param			role_value	query	string	true	"值"
//	@Param			role_name	query	string	false	"角色名称"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/role [post]
func AddRole(ctx *gin.Context) {
	role_value := ctx.Query("role_value")
	role_name := ctx.Query("role_name")
	result := models.AddRole(&models.Role{
		RoleValue: role_value,
		RoleName:  role_name,
	})
	ctx.JSON(http.StatusOK, result)
}

//	@Title			DeleteRole
//	@Summary		删除角色
//	@Description	删除角色
//	@Param			id	path	int	true	"角色id"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/role/{id} [delete]
func DeleteRole(ctx *gin.Context) {
	id := ctx.Param("id")
	post_id := utils.Str2Int(id, 0)
	result := models.DeleteRole(int(post_id))
	ctx.JSON(http.StatusOK, result)
}

//	@Title			GetCapabilities
//	@Summary		获取能力（权限）
//	@Description	获取能力（权限）
//	@Param			page	query	int	false	"页码"
//	@Param			size	query	int	false	"每页数量"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/cap [get]
func GetCapabilities(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageInt := utils.Str2Int(page)
	sizeInt := utils.Str2Int(size)
	users, count := models.GetCapabilities(int(pageInt), int(sizeInt))
	data := make(map[string]interface{})
	data["list"] = users
	data["count"] = count
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: data}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			GetCapability
//	@Summary		通过id获取能力（权限）
//	@Description	通过id获取能力（权限）
//	@Param			id	path	int	true	"能力id"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/cap/{id} [get]
func GetCapability(ctx *gin.Context) {
	id := ctx.Param("id")
	post_id := utils.Str2Int(id, 0)
	var result = utils.Result{}
	if post_id > 0 {
		post := models.GetCapability(int(post_id))
		if post.Valid() {
			result.ResultCode = utils.SUCCESS
			result.Data = post
		} else {
			result.ResultCode = utils.ERR_POST_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			AddCapability
//	@Summary		添加能力（权限）
//	@Description	添加能力（权限）
//	@Param			cap_value	query	string	true	"值"
//	@Param			cap_name	query	string	false	"角色名称"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/cap [post]
func AddCapability(ctx *gin.Context) {
	cap_value := ctx.Query("cap_value")
	cap_name := ctx.Query("cap_name")
	result := models.AddCapability(&models.Capability{
		CapValue: cap_value,
		CapName:  cap_name,
	})
	ctx.JSON(http.StatusOK, result)
}

//	@Title			DeleteCapability
//	@Summary		删除能力（权限）
//	@Description	删除能力（权限）
//	@Param			id	path	int	true	"能力id"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/cap/{id} [delete]
func DeleteCapability(ctx *gin.Context) {
	id := ctx.Param("id")
	post_id := utils.Str2Int(id, 0)
	result := models.DeleteCapability(int(post_id))
	ctx.JSON(http.StatusOK, result)
}

//	@Title			AddCapability2Role
//	@Summary		添加能力（权限）到角色
//	@Description	添加能力（权限）到角色
//	@Param			id		path	int	true	"角色id / role_value"
//	@Param			cap_id	query	int	true	"能力id / cap_value"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/role/{id}/cap [post]
func AddCapability2Role(ctx *gin.Context) {
	role_id := utils.Str2Int(ctx.Param("id"), 0)
	cap_id := utils.Str2Int(ctx.Query("cap_id"), 0)
	result := models.AddCapability2Role(uint(cap_id), uint(role_id))
	ctx.JSON(http.StatusOK, result)
}

//	@Title			DeleteCapability2Role
//	@Summary		从角色移除能力（权限）
//	@Description	从角色移除能力（权限）
//	@Param			id		path	int	true	"角色id"
//	@Param			cap_id	query	int	true	"能力id"
//	@Tags			角色与权限相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/role/{id}/cap [delete]
func DeleteCapability2Role(ctx *gin.Context) {
	role_id := utils.Str2Int(ctx.Param("id"), 0)
	cap_id := utils.Str2Int(ctx.Query("cap_id"), 0)
	result := models.DeleteCapabilityFromRole(uint(cap_id), uint(role_id))
	ctx.JSON(http.StatusOK, result)
}
