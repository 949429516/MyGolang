package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./new_book.html", "./book_list.html")
	v1 := r.Group("book")
	{
		v1.GET("", bookshow)
		// v1.GET()
		// v1.POST()
	}
	r.Run(":8000")
}
