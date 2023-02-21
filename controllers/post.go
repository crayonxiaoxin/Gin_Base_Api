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
	post_id, err := strconv.ParseInt(id, 0, 0)
	var result = utils.Result{}
	if err == nil && post_id != 0 {
		post := models.GetPost(int(post_id))
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
	post_id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		post_id = 0
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
	post.ID = uint(post_id)
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
	post_id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		post_id = 0
	}
	result := models.DeletePost(int(post_id))
	ctx.JSON(http.StatusOK, result)
}

//	@Title			Get PostMetas
//	@Summary		通过 post_id 获取元数据
//	@Description	通过 post_id 获取元数据
//	@Param			id			path	int		true	"post id"
//	@Param			meta_key	query	string	false	"Key，如果填写，只返回对应值"
//	@Tags			文章相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/posts/{id}/meta [get]
func GetPostMetas(ctx *gin.Context) {
	id := ctx.Param("id")
	meta_key := ctx.Query("meta_key")
	post_id, err := strconv.ParseInt(id, 0, 0)
	var result = utils.Result{}
	if err == nil && post_id != 0 {
		if len(meta_key) > 0 { // 如果有 meta_key，则获取单个值
			m := make(map[string]string)
			m[meta_key] = models.GetPostMetaValue(&models.PostMeta{PostId: uint(post_id), MetaKey: meta_key})
			result.Data = m
		} else { // 否则，获取所有相关值
			m := models.GetPostMetas(post_id)
			result.Data = m
		}
		result.ResultCode = utils.SUCCESS
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	ctx.JSON(http.StatusOK, result)
}

//	@Title			Update PostMeta
//	@Summary		新增或更新文章元数据
//	@Description	新增或更新文章元数据
//	@Param			id			path	int		true	"post id"
//	@Param			meta_key	query	string	true	"Key"
//	@Param			meta_value	query	string	false	"Value"
//	@Tags			文章相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/posts/{id}/meta [post]
func UpdatePostMeta(ctx *gin.Context) {
	id := ctx.Param("id")
	post_id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		post_id = 0
	}
	meta_key := ctx.Query("meta_key")
	meta_value := ctx.Query("meta_value")
	meta := models.PostMeta{PostId: uint(post_id), MetaKey: meta_key, MetaValue: meta_value}
	result := models.UpdatePostMeta(&meta)
	ctx.JSON(http.StatusOK, result)
}

//	@Title			DeletePostMeta
//	@Summary		删除文章元数据
//	@Description	删除文章元数据
//	@Param			id			path	int		true	"文章id"
//	@Param			meta_key	query	string	true	"Key"
//	@Tags			文章相关
//	@security		JwtAuth
//	@Success		200	{object}	utils.Result
//	@router			/posts/{id}/meta [delete]
func DeletePostMeta(ctx *gin.Context) {
	id := ctx.Param("id")
	meta_key := ctx.Query("meta_key")
	post_id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		post_id = 0
	}
	result := models.DeletePostMeta(uint(post_id), meta_key)
	ctx.JSON(http.StatusOK, result)
}
