package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tinolebat/golang/entity"
	"github.com/tinolebat/golang/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctxt *gin.Context) entity.Video
}

type Controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &Controller{
		service: service,
	}
}

func (c *Controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *Controller) Save(ctxt *gin.Context) entity.Video {
	var video entity.Video
	ctxt.ShouldBindJSON(&video)
	// ctxt.BindJSON(&video)
	c.service.Save(video)
	return video
}
