package db

import (
	"wklwukailun.com/blogger/model"
)

// 插入文章
func InsertAritcle(article *model.ArticleDetail) (articleId int64, err error) {
	// 验证
	if article == nil {
		return
	}
	sqlStr := `insert into article (content,summary,title,username,category_id,view_count,comment_count) values (?,?,?,?,?,?,?);`
	result, err := DB.Exec(sqlStr, article.Content, article.Summary, article.Title, article.Username,
		article.ArticleInfo.CategoryId, article.ViewCount, article.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

// 获取文章列表，作分页
func GetAricleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum <= 0 || pageSize <= 0 {
		return
	}
	// 时间降序排序
	sqlStr := `select id,category_id,summary,title,view_count,create_time,comment_count,username from article where
				status = 1 order by create_time desc limit ?,?;`
	err = DB.Select(&articleList, sqlStr, pageNum, pageSize)
	return
}

// 根据文章id，查询详细文章
func GetAricleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		return
	}
	sqlStr := `select id,category_id,summary,title,view_count,create_time,comment_count,username,content from article where id = ? and status = 1;`
	err = DB.Get(&articleDetail, sqlStr, articleId)
	return
}

// 根据分类id，查询这一类文章
func GetAricleListByCatrgoryId(articleId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum <= 0 || pageSize <= 0 {
		return
	}
	sqlStr := `select id,category_id,summary,title,view_count,create_time,comment_count,username from article where
	status = 1 and category_id = ? order by create_time desc limit ?,?;`
	err = DB.Select(&articleList, sqlStr, articleId, pageNum, pageSize)
	return
}
