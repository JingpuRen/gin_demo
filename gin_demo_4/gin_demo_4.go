package main

// todo 通过gin框架实现文件上传

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/web", func(context *gin.Context) {

	})

	r.Run("localhost:8080")
}
