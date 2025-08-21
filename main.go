package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default() //携带基础中间件启动
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
