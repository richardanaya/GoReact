package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	r := gin.Default()
	m := melody.New()
	r.GET("/api/ping", pingGet)
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	r.Static("/static", "./static")
	r.NoRoute(func(c *gin.Context) {
		c.File("./client/index.html")
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

func pingGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pongu",
	})
}
