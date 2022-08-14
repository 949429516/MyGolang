package service

import (
	"wklwukailun.com/blogger/dao/db"
	"wklwukailun.com/blogger/model"
)

// 获取文章和对应的分类
func GetArticleRecoedList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章那个列表
	articleInfoList, err := db.GetAricleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2.获取文章对应的分类(多个)
	ids := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(ids)
	if err != nil {
		return
	}
	// 返回页面做聚合
	// 遍历文章
	for _, article := range articleInfoList {
		// 根据文章生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		// 文章取出分类id
		categoryId := article.CategoryId
		// 遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// 根据多个文章的id，获取多个分类的id
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
	type a struct{}
	m := make(map[int64]a)
	// 遍历文章，得到每个id
	for _, article := range articleInfoList {
		// 从当前取出分类id
		categoryId := article.CategoryId
		// 去重复
		if _, ok := m[categoryId]; !ok {
			ids = append(ids, categoryId)
			m[categoryId] = a{}
		}
	}
	return
}

// 根据分类id获取文章和分类信息
func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.GetAricleListByCatrgoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	// 2.获取文章对应的分类(多个)
	ids := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(ids)
	if err != nil {
		return
	}
	// 返回页面做聚合
	// 遍历文章
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}
