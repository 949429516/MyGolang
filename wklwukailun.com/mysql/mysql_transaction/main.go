package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	// 数据库信息
	dsn := "root:19950811@tcp(127.0.0.1:3306)/sql_test"
	// 连接
	db, err = sql.Open("mysql", dsn) //这里不会校验用户名与密码正确，只会判断前面参数格式
	if err != nil {
		return err
	}
	err = db.Ping() //这里可以判断连接
	if err != nil {
		return err
	}
	//设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	//设置数据库闲置连接数
	db.SetConnMaxIdleTime(5)
	return
}
func transacrtion() {
	//开启事物
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 执行多个sql
	sqlStr1 := `update user set age = age -2 where id = 2;`
	sqlStr2 := `update user set age = age +2 where id = 1;`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("sqlstr1执行失败,回滚")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("sqlstr12执行失败,回滚")
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("commit执行失败,回滚")
		return
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	transacrtion()
}
