package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tinolebat/golang/controller"
	middleware "github.com/tinolebat/golang/custom-middleware"
	"github.com/tinolebat/golang/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func OutputLogFile() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {
	// server := gin.Default() => it intergrates the default Recovery and Logger middleware

	OutputLogFile()

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	server.Run(":8000")
}
