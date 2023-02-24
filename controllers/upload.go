package controllers

import (
	"hello_gin_api/models"
	"hello_gin_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Title			Upload
//	@Summary		上传
//	@Description	上传
//	@Param			file	formData	file	true	"文件"
//	@Tags			媒体相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/upload [post]
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

//	@Title			Delete
//	@Summary		删除文件
//	@Description	删除文件
//	@Param			id	path	int	true	"文件ID"
//	@Tags			媒体相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/media/{id} [delete]
func RemoveFile(ctx *gin.Context) {
	id := ctx.Param("id")
	post_id := utils.Str2Int(id, 0)
	result := models.DeleteMedia(int(post_id))
	ctx.JSON(http.StatusOK, result)
}

//	@Title			GetAll
//	@Summary		获取所有媒体
//	@Description	获取所有媒体
//	@Param			page	query	int		false	"页码"
//	@Param			size	query	int		false	"每页数量"
//	@Param			keyword	query	string	false	"关键词，默认空"
//	@Param			order	query	string	false	"默认：id desc"
//	@Param			uid		query	int		false	"要筛选的用户id"
//	@Tags			媒体相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/media [get]
func GetAllFiles(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	keyword := ctx.Query("keyword")
	order := ctx.Query("order")
	uid := ctx.Query("uid")
	options := &models.MediaListOptions{
		ListOptions: utils.ListOptions{
			Page:     utils.Str2Int(page),
			PageSize: utils.Str2Int(size),
			Keyword:  keyword,
			Order:    order,
		},
		Uid: utils.Str2Int(uid),
	}
	users, count := models.GetAllMedia(options)
	data := make(map[string]interface{})
	data["list"] = users
	data["count"] = count
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: data}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			Get
//	@Summary		通过id获取文件
//	@Description	通过id获取文件
//	@Param			id	path	int	true	"文件id"
//	@Tags			媒体相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/media/{id} [get]
func GetFile(ctx *gin.Context) {
	id := ctx.Param("id")
	post_id := utils.Str2Int(id, 0)
	var result = utils.Result{}
	if post_id > 0 {
		user := models.GetMedia(int(post_id))
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
