package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/get", func(c *gin.Context) {
			c.String(http.StatusOK, c.FullPath())
		})
	}
	r.Run()
}
