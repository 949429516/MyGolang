package db

import (
	"testing"
	"time"

	"wklwukailun.com/blogger/model"
)

func init() {
	// parseTime=true 将mysql中的时间类型自动解析为go结构体中的时间类型,不加会报错
	dns := "root:19950811@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 插入测试文章
func TestInsertAritcle(t *testing.T) {
	// 构建对象
	article := model.ArticleDetail{
		ArticleInfo: model.ArticleInfo{1, 1, "sss", "《学做人》", 1, time.Now(), 1, "wsb"},
		// 文章内容
		Content: "人在做天在看",
	}
	articleId, err := InsertAritcle(&article)
	if err != nil {
		return
	}
	t.Logf("articleId:%d\n", articleId)
}

func TestGetAricleList(t *testing.T) {
	articleList, err := GetAricleList(0, 2)
	if err != nil {
		return
	}
	t.Logf("article:%v", len(articleList))
}

func TestGetAricleDetail(t *testing.T) {
	articleList, err := GetAricleDetail(1)
	if err != nil {
		return
	}
	t.Logf("article:%v", articleList)
}
