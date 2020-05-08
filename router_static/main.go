package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/public", http.Dir("public"))
	r.StaticFile("/public.html", "./public/index.html")
	r.Run()
}
