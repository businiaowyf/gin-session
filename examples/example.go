package main

import (
	"log"
	"strconv"
	"time"

	"github.com/businiaowyf/gin-session/redisess"
	"github.com/gin-gonic/gin"
)

func main() {
	redisess.Init("session_key", 10, "127.0.0.1:6379", "")
	r := gin.Default()
	r.Use(redisess.NewSessionMiddleWare())

	r.GET("/test", func(c *gin.Context) {
		session, _ := redisess.SessStore.Get(c.Request, redisess.SessName)
		lastTime := session.Values["update_time"]
		session.Values["update_time"] = strconv.FormatInt(time.Now().Unix(), 10)
		session.Save(c.Request, c.Writer)
		log.Printf("Session ID=%v, lastTime=%v", session.ID, lastTime)
	})

	r.Run(":8080")
}
