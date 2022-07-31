package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		// 设置变量到Context的key中，可以通过get()获取
		c.Set("request", "中间件")
		// 执行中间件
		c.Next()
		// 执行后中间件做的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time", t2)
	}
}
func main() {
	r := gin.Default()
	// 注册中间件 全局中间件
	r.Use(MiddleWare())
	// {}为了代码规范
	{
		r.GET("/middleware", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面展示
			c.JSON(200, gin.H{"request": req})
		})
		// 根路由定义的是局部中间件
		r.GET("/middleware2", MiddleWare(), func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面展示
			c.JSON(200, gin.H{"request": req})
		})
	}
	r.Run(":8000")
}
