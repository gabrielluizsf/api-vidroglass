package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/mariarobertap/api-vidroglass/models"
    "github.com/mariarobertap/api-vidroglass/service"

)


type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context) models.Video
}

type controllerVideo struct {
	service service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	return &controllerVideo {
		service: service,
	}
}

func (c *controllerVideo) FindAll() []models.Video {
	return c.service.FindAll()
}


func (c *controllerVideo) Save(ctx *gin.Context) models.Video{
	var video models.Video
	ctx.BindJSON(&video)
	c.service.Save(video)

	return video
}