package service

import (
	"wklwukailun.com/blogger/dao/db"
	"wklwukailun.com/blogger/model"
)

// 获取所有分类
func GetALLCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}
