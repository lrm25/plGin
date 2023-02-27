package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed public/*
var fs embed.FS

func main() {
	router := gin.Default()

	router.GET("/helloworld/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})
	router.POST("/helloworld/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Word POST call")
	})
	router.StaticFile("/", "./public/test.html")
	router.Static("/public", "./public")
	router.StaticFS("/fs", http.FileSystem(http.FS(fs)))
	router.Run(":3000")
}
