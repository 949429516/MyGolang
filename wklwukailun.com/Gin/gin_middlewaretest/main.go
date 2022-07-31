package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	//统计
	since := time.Since(start)
	fmt.Println("程序用时:", since)
}
func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)

}
func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)

}
func main() {
	r := gin.Default()
	// 注册中间件 全局中间件
	r.Use(myTime)
	shoppingGroup := r.Group("/shopping")
	// {}为了代码规范
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run(":8000")
}
