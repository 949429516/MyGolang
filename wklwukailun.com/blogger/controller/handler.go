package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"wklwukailun.com/blogger/service"
)

// 访问主页的控制器
func IndexHandle(c *gin.Context) {
	// 从service取数据
	// 1.加载文章数据
	articleRecordList, err := service.GetArticleRecoedList(0, 15)
	if err != nil {
		// c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		c.JSON(http.StatusBadRequest, gin.H{"res": "500"})

		return
	}
	// 2.加载分类数据
	categoryList, err := service.GetALLCategoryList()
	if err != nil {
		// c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		c.JSON(http.StatusBadRequest, gin.H{"res": "500"})

		return
	}
	// gin.H本质是一个map
	// var data map[string]interface{} = make(map[string]interface{}, 16)
	// data["article_list"] = articleRecordList
	// data["category_list"] = categoryList
	// c.HTML(http.StatusOK, "views/index.html",data)
	// c.HTML(http.StatusOK, "views/index.html", gin.H{
	// 	"article_list":  articleRecordList,
	// 	"category_list": categoryList,
	// })
	c.JSON(http.StatusOK, gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

// 点击分类云进行分类
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	// 转换int64
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		//c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		c.JSON(http.StatusBadRequest, gin.H{"res": "500"})
		return
	}
	// 根据分类获取文章列表
	articleRecordList, err := service.GetArticleRecordListById(int(categoryId), 0, 15)
	if err != nil {
		// c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		c.JSON(http.StatusBadRequest, gin.H{"res": "500"})
		return
	}
	// 2.加载分类数据
	categoryList, err := service.GetALLCategoryList()
	if err != nil {
		// c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		c.JSON(http.StatusBadRequest, gin.H{"res": "500"})
		return
	}
	// c.HTML(http.StatusOK, "views/index.html", gin.H{
	// 	"article_list":  articleRecordList,
	// 	"category_list": categoryList,
	// })
	c.JSON(http.StatusOK, gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})

}
