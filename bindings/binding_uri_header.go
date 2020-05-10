package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Person2 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person2
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	r.GET("/", func(c *gin.Context) {
		h := testHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
			return
		}

		fmt.Printf("%#v\n", h)
		c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})
	r.Run()
}
