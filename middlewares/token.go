package middlewares

import (
	"hello_gin_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 验证 token
func NeedToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := utils.Result{}
		tokenString := ctx.GetHeader("token")
		rc, _ := utils.ParseToken(tokenString)
		if !rc.Success() {
			result.ResultCode = rc
			ctx.JSON(http.StatusOK, result)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
