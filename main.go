package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

//type PostParams struct {
//	Name string `json:"name" uri:"name"`
//	Age  int    `json:"age" uri:"age"`
//	Sex  bool   `json:"sex" uri:"sex"`
//}

//type PostParams struct {
//	Name string `json:"name"`
//	Age  int    `json:"age" binding:"required,mustBig"` //mustBig 自定义匹配规则 required,mustBig中间不能有空格 shift 查了好久
//	Sex  bool   `json:"sex"`
//}
//
//func mustBig(f1 validator.FieldLevel) bool {
//
//	if f1.Field().Interface().(int) <= 18 {
//		return false
//	}
//	//fmt.Println(f1.Field().Interface().(int))
//	return true
//}

func main() {
	r := gin.Default()
	r.POST("/testUpload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//c.SaveUploadedFile(file, "./upload/"+file.Filename)
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./upload/" + file.Filename)
		defer out.Close()
		io.Copy(out, in)

		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
		c.File("./upload/" + file.Filename)

		//c.JSON(200, gin.H{
		//	"msg":    file,
		//	"status": "ok",
		//})
	})

	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	e := v.RegisterValidation("mustBig", mustBig)
	//	fmt.Println(e)
	//}
	//
	//r.POST("/testBind", func(c *gin.Context) {
	//	var p PostParams
	//	err := c.ShouldBindJSON(&p)
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"msg":  "失败",
	//			"data": gin.H{},
	//		})
	//	} else {
	//		c.JSON(200, gin.H{
	//			"msg":  "success",
	//			"data": p,
	//		})
	//	}
	//})

	//r.POST("/testBind", func(c *gin.Context) {
	//	var p PostParams
	//	err := c.ShouldBindJSON(&p)
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"msg":  "失败",
	//			"data": gin.H{},
	//		})
	//	} else {
	//		c.JSON(200, gin.H{
	//			"msg":  "success",
	//			"data": p,
	//		})
	//	}
	//})

	//r := gin.Default() //携带基础中间件启动
	//r.GET("/path/:id", func(c *gin.Context) {
	//	id := c.Param("id") //id 占位
	//	user := c.DefaultQuery("user", "admin")
	//	pwd := c.Query("pwd")
	//	c.JSON(200, gin.H{
	//		"id":   id,
	//		"user": user,
	//		"pwd":  pwd,
	//	})
	//})
	//r.POST("/path", func(c *gin.Context) {
	//	user := c.DefaultPostForm("user", "admin")
	//	pwd := c.PostForm("pwd")
	//	c.JSON(200, gin.H{
	//		"user": user,
	//		"pwd":  pwd,
	//	})
	//})
	//r.DELETE("/path/:id", func(c *gin.Context) {
	//	id := c.Param("id") //id 占位
	//	c.JSON(200, gin.H{
	//		"id": id,
	//	})
	//})
	//r.PUT("/path", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"status":  "success",
	//		"message": "pong",
	//	})
	//})
	r.Run(":14005") // listen and serve on 0.0.0.0:8080
}
