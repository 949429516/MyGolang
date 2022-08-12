package db

import "testing"

func init() {
	// parseTime=true 将mysql中的时间类型自动解析为go结构体中的时间类型,不加会报错
	dns := "root:19950811@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 获取单个分类
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

// 获取多个分类
func TestGetCategoryList(t *testing.T) {
	categoryList, err := GetCategoryList([]int64{1, 2, 3})
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d,category:%#v\n", v.CategoryId, v)
	}
}

// 获取所有分类
func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d,category:%#v\n", v.CategoryId, v)
	}
}
