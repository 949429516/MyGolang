package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

// 数据库初始化
func Init(dns string) (err error) {
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return
	}
	// 查看连接成功，测试代码正式上线注解
	err = DB.Ping()
	if err != nil {
		return
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return
}

// 分类相关操作（添加、查询、查1个分类、查多个分类、查看所有分类）

// 文章相关操作(添加文章、查询所有文章、根据文章id查看内容)
