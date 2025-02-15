package main

// gin的helloWorld

func main() {
	/**
	todo 返回JSON格式代码
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"name": "jingpu", "password": "daohaozhe250"})
	})
	r.Run("localhost:8080")
	*/

	/**
	todo 获取请求中的参数
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		name := c.Query("query") // todo 通过c.Query()方法获取请求中"query"字段的参数
	// c.DefaultQuery("query","somebody") todo 如果"query"字段有值，则用其值；否则我们就默认是"somebody"
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run("localhost:8080")
	*/

	/**
	todo 获取前端提交的表单中的数据
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// todo 通过c.PostForm()方法获取表单中的数据
		name := c.PostForm("name")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"name": name, "password": password,
		})
	})
	r.Run("localhost:8080")
	*/

}
