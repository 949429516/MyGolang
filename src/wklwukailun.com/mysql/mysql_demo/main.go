package main

//go get -u github.com/go-sql-driver/mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //init
)

//连接mysql
func main() {
	// 数据库信息
	dsn := "root:19950811@tcp(127.0.0.1:3306)/goday10"
	// 连接
	db, err := sql.Open("mysql", dsn) //这里不会校验用户名与密码正确，只会判断前面参数格式
	if err != nil {
		panic(err)
	}
	err = db.Ping() //这里可以判断连接
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}
