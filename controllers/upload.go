package controllers

import (
	"hello_gin_api/models"
	"hello_gin_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Title       Upload
// @Summary     上传
// @Description 上传
// @Param       token header   string true "token"
// @Param       file  formData file   true "文件"
// @Tags        媒体相关
// @Success     200 {object} utils.Result
// @router      /upload [post]
func UploadFile(ctx *gin.Context) {
	// 文件
	fh, err := ctx.FormFile("file")
	result := utils.Result{}
	if err != nil {
		result.ResultCode = utils.ERR_PARAMS
	} else {
		tokenString := ctx.GetHeader("token")
		uid := utils.GetUidFromToken(tokenString)
		result = models.UploadMedia(fh, uid)
	}
	ctx.JSON(http.StatusOK, result)
}

// @Title       Delete
// @Summary     删除文件
// @Description 删除文件
// @Param       token header string true "token"
// @Param       id    path   int    true "文件ID"
// @Tags        媒体相关
// @Success     200 {object} utils.Result
// @router      /media/{id} [delete]
func RemoveFile(ctx *gin.Context) {
	id := ctx.Param("id")
	mid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		mid = 0
	}
	result := models.DeleteMedia(int(mid))
	ctx.JSON(http.StatusOK, result)
}

// @Title       GetAll
// @Summary     获取所有媒体
// @Description 获取所有媒体
// @Param       token header string true  "登入后返回的token"
// @Param       page  query  int    false "页码"
// @Param       size  query  int    false "每页数量"
// @Tags        媒体相关
// @Success     200 {object} utils.Result
// @router      /media [get]
func GetAllFiles(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageInt := utils.Str2Int(page)
	sizeInt := utils.Str2Int(size)
	users, count := models.GetAllMedia(int(pageInt), int(sizeInt))
	data := make(map[string]interface{})
	data["list"] = users
	data["count"] = count
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: data}
	ctx.JSON(http.StatusOK, result)
}

// @Title       Get
// @Summary     通过id获取文件
// @Description 通过id获取文件
// @Param       token header string true "登入后返回的token"
// @Param       id    path   int    true "The key for staticblock"
// @Tags        媒体相关
// @Success     200 {object} utils.Result
// @router      /media/{id} [get]
func GetFile(ctx *gin.Context) {
	id := ctx.Param("id")
	mid, err := strconv.ParseInt(id, 0, 0)
	var result = utils.Result{}
	if err == nil && mid != 0 {
		user := models.GetMedia(int(mid))
		if user.Valid() {
			result.ResultCode = utils.SUCCESS
			result.Data = user
		} else {
			result.ResultCode = utils.ERR_UPLOAD_FILE_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	ctx.JSON(http.StatusOK, result)
}
