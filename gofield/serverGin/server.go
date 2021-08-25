package main

import (
	"github.com/gin-gonic/gin"
	"github.com/testzone/gofield/serverGin/controller"
	"github.com/testzone/gofield/serverGin/service"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main()  {
	server:= gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "OK!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
	ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	_=server.Run(":8080")
}