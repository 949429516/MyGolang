package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func bookshow(c *gin.Context) {
	title := queryRowMore()
	c.HTML(200, "book_list.html", gin.H{"data": title})
}
func addshow(c *gin.Context) {
	c.HTML(200, "new_book.html", gin.H{})
}
func add(c *gin.Context) {
	title := c.PostForm("title")
	price := c.PostForm("price")
	priceint, _ := strconv.Atoi(price)
	err := insert(title, priceint)
	if err == nil {
		c.Redirect(http.StatusMovedPermanently, "/book")
	} else {
		c.String(http.StatusBadRequest, "add book failed")
	}
}
func del(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	idint, _ := strconv.Atoi(id)
	err := delete(idint)
	if err == nil {
		c.Redirect(http.StatusMovedPermanently, "/book")
	} else {
		c.String(http.StatusBadRequest, "delete failed")
	}
}
