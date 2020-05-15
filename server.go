package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tinolebat/golang/controller"
	"github.com/tinolebat/golang/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))

	})

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())

	})

	server.Run(":8000")
}
