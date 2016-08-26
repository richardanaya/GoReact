package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api/ping", pingGet)
	r.Static("/static", "./static")
	r.NoRoute(func(c *gin.Context) {
		c.File("./index.html")
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

func pingGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pongu",
	})
}
