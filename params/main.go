package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.GET("/users/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, fmt.Sprintf("your name is %s", name))
	//})

	r.GET("/users/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := "name:" + name + ";action:" + action
		c.String(http.StatusOK, message)
	})

	r.POST("/users/:name/*action", func(c *gin.Context) {
		c.String(http.StatusOK, c.FullPath())
	})

	r.GET("/query", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstName", "Sn")
		lastName := c.Query("lastName")
		c.JSON(http.StatusOK, gin.H{
			"firstname": firstName,
			"lastname": lastName,
		})
	})

	/**
	 * POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	 * Content-Type: application/x-www-form-urlencoded
     *
	 * names[first]=thinkerou&names[second]=tianou&message=xxx
	 */
	r.POST("/query", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"nick":    nick,
			"ids":     ids,
			"names":   names,
		})
	})
	r.Run(":8080")
}
