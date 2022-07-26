package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// form表单接收
func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		// 表单设置默认参数
		type1 := c.DefaultPostForm("type", "alert")
		username := c.PostForm("username")
		password := c.PostForm("password")
		hobbys := c.PostFormArray("hobbys")
		c.String(http.StatusOK, fmt.Sprintf("type is %s,username is %s,password is %s,hobbys is %v\n", type1, username, password, hobbys))
	})
	r.Run(":8000")
}
