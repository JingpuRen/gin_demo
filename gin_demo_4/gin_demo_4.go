package main

// tip 通过gin框架实现http重定向和路由重定向

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// tip http请求重定向
	router.GET("/index", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})

	// tip 内部路由重定向
	router.GET("/web1", func(context *gin.Context) {
		context.Request.URL.Path = "/web2" // 修改请求中的URL信息
		router.HandleContext(context)      // 继续后续的处理
	})
	router.GET("./web2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "you are succeed",
		})
	})

	router.Run("localhost:8080")
}
