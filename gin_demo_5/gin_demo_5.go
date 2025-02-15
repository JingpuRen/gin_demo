package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// tip gin框架中的路由和路由组

func main() {
	router := gin.Default()

	// tip 1:常见的get/post/put/delete路由
	router.GET("/get", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "you are right",
		})
	})

	// tip 2:可以匹配所有的请求方法
	// tip 注意，这里的请求方法意思是前面get这些，并不是后面的那个路径
	router.Any("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "this is /test",
		})
	})

	// tip 3:如果请求方法不能匹配到任何的路由，那么就进入该路由
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"data": "你的请求不正确",
		})
	})

	// tip 4:路由组，当具有公共前缀的时候，我们一般会使用路由组
	userGroup := router.Group("/user")
	{
		userGroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": "this is get",
			})
		})
		userGroup.GET("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"data": "this is get2",
			})
		})

	}
	router.Run("localhost:8080")
}
