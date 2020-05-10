package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	r := gin.New()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s", c.DefaultQuery("name", "longerwu"))
		fmt.Println("continue...")
	})
	r.GET("/panic", func(c *gin.Context) {
		a,b := 1,0
		c.JSON(http.StatusOK, a/b)
		fmt.Println("continue...")
	})
	r.Run()
}
