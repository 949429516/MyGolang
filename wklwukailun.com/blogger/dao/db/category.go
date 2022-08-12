package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"wklwukailun.com/blogger/model"
)

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

// 获取多个分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlStr, args, err := sqlx.In(`select id,category_name,category_no from category where id in (?);`, categoryIds)
	if err != nil {
		return
	}
	fmt.Println(sqlStr, args)
	// 查询
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

// 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := `select id,category_name,category_no from category order by create_time asc;`
	err = DB.Select(&categoryList, sqlStr)
	return
}
