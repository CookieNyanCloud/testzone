package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/testzone/gofield/serverGin/entity"
	"github.com/testzone/gofield/serverGin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController  {
	return &controller{
		service,
	}
}

func (c *controller) FindAll() []entity.Video  {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video  {
	var video entity.Video
	_=ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}