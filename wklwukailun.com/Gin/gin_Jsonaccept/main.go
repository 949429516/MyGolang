package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// JSON绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 声明变量
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名
		if json.Username != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}

// func main() {
// 	// 1.创建路由
// 	r := gin.Default()
// 	// JSON绑定
// 	r.GET("/loginJSON/:username/:password", func(c *gin.Context) {
// 		// 声明变量
// 		var login Login
// 		if err := c.ShouldBindUri(&login); err != nil {
// 			// 返回错误信息
// 			// gin.H封装了生成json数据工具
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		// 判断用户名
// 		if login.Username != "root" || login.Password != "admin" {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"status": "200"})
// 	})
// 	r.Run(":8000")
// }
// func main() {
// 	// 1.创建路由
// 	r := gin.Default()
// 	// JSON绑定
// 	r.POST("loginJSON", func(c *gin.Context) {
// 		// 声明变量
// 		var form Login
// 		if err := c.Bind(&form); err != nil { // 以form形式提交POST请求
// 			// 返回错误信息
// 			// gin.H封装了生成json数据工具
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		// 判断用户名
// 		if form.Username != "root" || form.Password != "admin" {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"status": "200"})
// 	})
// 	r.Run(":8000")
// }
