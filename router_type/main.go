package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "GET")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "POST")
	})
	r.PUT("/put", putHandler)
	r.Handle(http.MethodDelete, "/delete", func(c *gin.Context) {
		c.String(200, "DELETE")
	})
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "ANY")
	})
	r.Run()
}

func putHandler(c *gin.Context) {
	c.String(200, "PUT")
}
