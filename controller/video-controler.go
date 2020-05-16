package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tinolebat/golang/entity"
	"github.com/tinolebat/golang/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctxt *gin.Context) error
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
func (c *Controller) Save(ctxt *gin.Context) error {
	var video entity.Video
	err := ctxt.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}

func Handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
