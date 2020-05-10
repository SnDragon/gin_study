package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

// IP白名单中间件
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		whiteIPList := []string{"127.0.0.1"}
		clientIP := c.ClientIP()
		passFlag := false
		for _, ip := range whiteIPList {
			if ip == clientIP {
				passFlag = true
				break
			}
		}
		if !passFlag {
			c.String(http.StatusUnauthorized, "%s not in ip whitelist", clientIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.New()
	r.Use(Logger(), IPAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
		c.JSON(http.StatusOK,"success")
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
