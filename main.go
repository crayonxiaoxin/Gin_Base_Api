package main

import (
	"hello_gin_api/controllers"
	"hello_gin_api/docs"
	"hello_gin_api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@Version					v1
//	@Title						API Demo
//	@Description				基于 golang 构建的 API 项目
//	@BasePath					/v1
//	@contact.name				Code Resources
//	@contact.url				https://github.com/crayonxiaoxin/Gin_Base_Api
//	@securityDefinitions.apikey	JwtAuth
//	@in							header
//	@name						token

func main() {

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/v1"

	// 最大上传大小 8M
	r.MaxMultipartMemory = 8 << 20

	v1 := r.Group("/v1")
	{
		v1.POST("/login", controllers.Login)
		v1.POST("/register", controllers.Register)

		// 设置中间件：验证token
		v1.Use(middlewares.NeedToken())
		{
			user := v1.Group("/user")
			{
				user.GET("/", controllers.GetUsers)
				user.GET("/:id", controllers.GetUser)
				user.POST("/", controllers.AddUser)
				user.PUT("/:id", controllers.UpdateUser)
				user.DELETE("/:id", controllers.DeleteUser)

				// 元数据
				user.GET("/:id/meta", controllers.GetUserMetas)
				user.POST("/:id/meta", controllers.UpdateUserMeta)
			}

			// 文件上传
			v1.POST("/upload", controllers.UploadFile)
			// 媒体
			media := v1.Group("/media")
			{
				media.GET("/", controllers.GetAllFiles)
				media.GET("/:id", controllers.GetFile)
				media.DELETE("/:id", controllers.RemoveFile)
			}

			// 文章
			posts := v1.Group("/posts")
			{
				posts.GET("/", controllers.GetPosts)
				posts.GET("/:id", controllers.GetPost)
				posts.POST("/", controllers.AddPost)
				posts.PUT("/:id", controllers.UpdatePost)
				posts.DELETE("/:id", controllers.DeletePost)
			}
		}
	}

	// swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 重定向至文档
	r.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})

	// 404
	r.NoRoute(controllers.Error404)

	// 设置静态文件路径
	r.Static("/uploads", "uploads")

	r.Run(":8083")
}
