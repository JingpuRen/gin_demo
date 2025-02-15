package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin获取URI路径参数

func main() {
	// 获取默认的路由引擎
	r := gin.Default()

	// 定义方法
	r.GET("/:username/:age", func(c *gin.Context) {
		username := c.Param("username")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": username,
			"age":  age,
		})
	})

	// todo 如果有多个GET方法的时候，前面一定要有前缀，例如/user/:username/:password

	// 指定服务的ip地址和端口
	r.Run("localhost:8080")
}
