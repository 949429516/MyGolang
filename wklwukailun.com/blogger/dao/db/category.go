package db

import "wklwukailun.com/blogger/model"

// 添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	// 分类相关操作（添加、查询、查1个分类、查多个分类、查看所有分类）
	sqlStr := `insert into category(category_name,category_no) values (?,?);`
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

// 获取单个分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := `select id,category_name,category_no from category where id = ?;`
	err = DB.Get(category, sqlStr, id)
	return
}
