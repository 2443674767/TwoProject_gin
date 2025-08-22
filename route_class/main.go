package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 中间件
func middel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("first1")
		c.Next()
		fmt.Println("second1")
	}
}

func middeltwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("first2")
		c.Next()
		fmt.Println("second2")
	}
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1").Use(middel()).Use(middeltwo()) //路由管理 路由分块
	v1.GET("test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
		fmt.Println("分组")
	})
	r.Run(":14006")
}
