package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 1.异步,在启动新的groutine时，不应该使用原始上下文，必须使用他的只读副本
	// 异步访问直接返回值
	r.GET("/long_async", func(context *gin.Context) {
		// 需要一个副本
		copyContext := context.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行:", copyContext.Request.URL.Path)
		}()
	})
	// 2.同步
	r.GET("/long_sync", func(context *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行:", context.Request.URL.Path)
	})
	r.Run(":8000")
}
