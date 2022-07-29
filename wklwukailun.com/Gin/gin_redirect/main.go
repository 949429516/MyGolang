package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/redirect", func(context *gin.Context) {
		// 支持内部和外部重定向
		context.Redirect(http.StatusMovedPermanently, "https://baidu.com")
	})
	r.Run(":8000")
}
