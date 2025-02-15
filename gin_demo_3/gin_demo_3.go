package main

// todo gin参数绑定，因为前端有可能传过来的是结构体，也就是对象

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserInfo todo 其实就是将前端传送过来的JSON数组解析成对应的结构体
type UserInfo struct {
	// todo 后面的`form:"name"`的意思是传递过来的表单数据中的"name"字段的值应该绑定到我们这里的"Name"上
	// todo 这样前端不管前端是以json格式传递过来还是以form表单格式传递过来我们都能够找到对应的变量的值
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/web", func(c *gin.Context) {
		var userInfo UserInfo
		// todo 我们需要将数据绑定到userInfo上，因此需要对传入的数据进行修改，因此我们要传入userInfo的地址
		err := c.ShouldBind(&userInfo)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "g",
			})
		} else {
			fmt.Println(userInfo.Name)
			fmt.Println(userInfo.Password)
			c.JSON(http.StatusOK, gin.H{
				"message": "you are succeed",
			})
		}
	})

	r.POST("/json", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.ShouldBind(&userInfo)
		if err != nil {
			context.JSON(http.StatusForbidden, gin.H{
				"message": "you are wrong",
			})
		} else {
			fmt.Println(userInfo.Name)
			fmt.Println(userInfo.Password)
			context.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "you are succeed",
			})
		}

	})

	r.Run("localhost:8080")
}
