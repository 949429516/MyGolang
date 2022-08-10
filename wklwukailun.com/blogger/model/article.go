package model

import "time"

// 定义文章结构体

type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

// 用于文章详情页的实体，为了提升效率把文章内容平时不加载,单独创建一个结构体嵌套
type ArticleDetail struct {
	ArticleInfo
	Category
	// 文章内容
	Content string `db:"content"`
}

// 用于文章上一篇与下一篇
type ArticleRecord struct {
	ArticleInfo
	Category
}
