package main

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func hello(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "hello",
  })
}

func upload(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "upload",
  })
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/hello", hello)
  router.POST("/upload", upload)
  router.GET(
    "user/:name/*action", func(ctx *gin.Context) {
      name:=ctx.Param("name")
      action:=ctx.Param("action")
      ctx.JSON(200, gin.H{
        "name": name,
        "action": action,
      })
    })
	router.Run(":8080") // 默认监听 0.0.0.0:8080
}
