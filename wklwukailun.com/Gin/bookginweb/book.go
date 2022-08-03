package main

import (
	"github.com/gin-gonic/gin"
)

func bookshow(c *gin.Context) {
	title := queryRowMore()
	c.HTML(200, "book_list.html", gin.H{"data": title})
}
