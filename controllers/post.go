package controllers

import (
	"hello_gin_api/models"
	"hello_gin_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Title			GetPosts
//	@Summary		获取所有文章
//	@Description	获取所有文章
//	@Param			page	query	int	false	"页码"
//	@Param			size	query	int	false	"每页数量"
//	@Tags			文章相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/posts [get]
func GetPosts(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageInt := utils.Str2Int(page)
	sizeInt := utils.Str2Int(size)
	posts, count := models.GetPosts(int(pageInt), int(sizeInt))
	data := make(map[string]interface{})
	data["list"] = posts
	data["count"] = count
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: data}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			GetPost
//	@Summary		通过id获取文章
//	@Description	通过id获取文章
//	@Param			id	path	int	true	"文章id"
//	@Tags			文章相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/posts/{id} [get]
func GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	pid, err := strconv.ParseInt(id, 0, 0)
	var result = utils.Result{}
	if err == nil && pid != 0 {
		post := models.GetPost(int(pid))
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

//	@Title			AddPost
//	@Summary		添加文章
//	@Description	添加文章
//	@Param			post_title		query	string	true	"标题"
//	@Param			post_content	query	string	true	"内容"
//	@Param			post_status		query	string	false	"状态：publish/draft"
//	@Param			post_type		query	string	false	"类型"
//	@Param			post_parent		query	string	false	"父文章"
//	@Tags			文章相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/posts [post]
func AddPost(ctx *gin.Context) {
	tokenString := ctx.GetHeader("token")
	uid := utils.GetUidFromToken(tokenString)
	post_title := ctx.Query("post_title")
	post_content := ctx.Query("post_content")
	post_status := ctx.Query("post_status")
	post_type := ctx.Query("post_type")
	post_parent := ctx.Query("post_parent")
	result := models.AddPost(&models.Post{
		Uid:         uint(uid),
		PostTitle:   post_title,
		PostContent: post_content,
		PostStatus:  post_status,
		PostType:    post_type,
		PostParent:  uint(utils.Str2Int(post_parent)),
	})
	ctx.JSON(http.StatusOK, result)
}

//	@Title			UpdatePost
//	@Summary		更新文章
//	@Description	更新文章
//	@Param			id				path	int		true	"文章id"
//	@Param			post_title		query	string	true	"标题"
//	@Param			post_content	query	string	true	"内容"
//	@Param			post_status		query	string	false	"状态：publish/draft"
//	@Param			post_type		query	string	false	"类型"
//	@Param			post_parent		query	string	false	"父文章"
//	@Tags			文章相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/posts/{id} [put]
func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	pid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		pid = 0
	}
	post_title := ctx.Query("post_title")
	post_content := ctx.Query("post_content")
	post_status := ctx.Query("post_status")
	post_type := ctx.Query("post_type")
	post_parent := ctx.Query("post_parent")
	post := models.Post{
		PostTitle:   post_title,
		PostContent: post_content,
		PostStatus:  post_status,
		PostType:    post_type,
		PostParent:  uint(utils.Str2Int(post_parent)),
	}
	post.ID = uint(pid)
	result := models.UpdatePost(&post)
	ctx.JSON(http.StatusOK, result)
}

//	@Title			DeletePost
//	@Summary		删除文章
//	@Description	删除文章
//	@Param			id	path	int	true	"文章id"
//	@Tags			文章相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/posts/{id} [delete]
func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		uid = 0
	}
	result := models.DeletePost(int(uid))
	ctx.JSON(http.StatusOK, result)
}
