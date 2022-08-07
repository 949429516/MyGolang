package main

import "github.com/gin-gonic/gin"

func main() {
	initDB()
	r := gin.Default()
	r.LoadHTMLFiles("./new_book.html", "./book_list.html")
	v1 := r.Group("book")
	{
		v1.GET("", bookshow)
		v1.GET("add", addshow)
		v1.POST("add", add)
		v1.GET("delete/", del)
	}
	r.Run(":8000")
}
