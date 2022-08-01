package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置cookie
			// name cookie键, value cookie值, maxAge 有效期单位秒, path cookie锁在目录, domain 域名, secure 是否能通过https访问, httpOnly 是否允许别人通过js获取cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "127.0.0.1", "./", false, true)
			fmt.Println(cookie)
		}
	})
	r.Run(":8000")
}
