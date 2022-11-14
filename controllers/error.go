package controllers

import (
	"hello_gin_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error404(ctx *gin.Context) {
	result := utils.Result{}
	result.ResultCode = utils.ERR_404
	ctx.JSON(http.StatusNotFound, result)
}
