package main

import (
	"hello_gin_api/controllers"
	"hello_gin_api/docs"
	"hello_gin_api/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Version     v1
// @Title       Test API
// @Description Gin API 基础工程
// @BasePath    /v1

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
				user.GET("/", controllers.GetAllUsers)
				user.GET("/:id", controllers.GetUser)
				user.POST("/:id", controllers.AddUser)
				user.PUT("/:id", controllers.UpdateUser)
				user.DELETE("/:id", controllers.DeleteUser)
			}

			// 文件上传
			v1.POST("/upload", controllers.UploadFile)

			media := v1.Group("/media")
			{
				media.GET("/", controllers.GetAllFiles)
				media.GET("/:id", controllers.GetFile)
				media.DELETE("/:id", controllers.RemoveFile)
			}
		}
	}

	// swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 404
	r.NoRoute(controllers.Error404)

	// 设置静态文件路径
	r.Static("/uploads", "uploads")

	r.Run(":8081")
}
