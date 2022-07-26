package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 上传文件
func main() {
	r := gin.Default()
	// 限制表单文件大小 8MB,默认32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 表单文件
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有文件
		files := form.File["file"]
		// 遍历所有文件
		for _, file := range files {
			// 逐个存储
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d", len(files)))
	})
	r.Run(":8000")
}
