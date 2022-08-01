package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过则不再进行后续处理
		c.Abort()
		return
	}
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./home.html", "./login.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != "root" || password != "root" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password error"})
		} else {
			c.SetCookie("abc", "123", 60, "127.0.0.1", "./", false, true)
			c.Redirect(http.StatusMovedPermanently, "/home")
		}
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{})
	})
	r.Run(":8000")
}
