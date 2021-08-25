package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/testzone/gofield/serverGin/controller"
	"github.com/testzone/gofield/serverGin/middlewares"
	"github.com/testzone/gofield/serverGin/service"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupOutput()  {
	f,_:=os.Create("/serverGin/gin.log")
	gin.DefaultWriter =io.MultiWriter(f,os.Stdout)
}

func main()  {
	setupOutput()

	server:= gin.New()

	server.Static("/css","./templates/css")

	server.LoadHTMLGlob("serverGin/templates/*.html")

	server.Use(gin.Recovery(),
		middlewares.Logger(),
		//middlewares.BasicAuth(),
		gindump.Dump(),
	)

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message": "OK!",
		})
	})

	apiRoutes:= server.Group("/api",middlewares.BasicAuth())
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
			//ctx.Header("X-Frame-Options", "")
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err:= videoController.Save(ctx)
			if err!=nil{
				ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK,gin.H{"message": "okay"})
			}
		})
	}
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos",videoController.ShowAll)
	}

	port:= os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	err:=server.Run(":"+port)
	if err== nil {
		fmt.Println("Listening on ", port)
	}

}