package main

import (
	"github.com/gin-gonic/gin"
	"wklwukailun.com/blogger/controller"
	"wklwukailun.com/blogger/dao/db"
)

func main() {
	router := gin.Default()
	dns := "root:19950811@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	// 加载静态文件
	//router.Static("/static/", "./static")
	// 加载模板
	//router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	_ = router.Run(":8000")
}
