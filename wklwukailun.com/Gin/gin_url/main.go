package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// http://127.0.0.1:8000/welcome/?name=sb
// url方式接收参数
func main() {
	r := gin.Default()
	r.GET("/welcome/", func(c *gin.Context) {
		//DefaultQuery第二个参数为默认值
		name := c.DefaultQuery("name", "jack")
		c.String(http.StatusOK, name)
	})
	r.Run(":8000")
}
