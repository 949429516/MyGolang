package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// http://127.0.0.1:8000/user/w/b
// ResAPI
func main() {
	// 1.创建路由
	// 默认使用了两个中间件Logger(),Recovery()
	r := gin.Default()
	// api参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.Run(":8000")
}
