package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Book struct {
	Id    int    `json:"Id"`
	Title string `json:"Title"`
	Price int    `json:"Price"`
}

func initDB() (err error) {
	// 数据库信息
	dsn := "root:19950811@tcp(127.0.0.1:3306)/goday10"
	// 连接
	db, err = sql.Open("mysql", dsn) //这里不会校验用户名与密码正确，只会判断前面参数格式
	if err != nil {
		return
	}
	err = db.Ping() //这里可以判断连接
	if err != nil {
		return
	}
	//设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	//设置数据库闲置连接数
	db.SetConnMaxIdleTime(5)
	return
}

// 查询多条
func queryRowMore() (data []Book) {
	sqlStr := `select id,title,price from book;`
	rows, err := db.Query(sqlStr)
	if err != nil {
		return
	}
	//一定要关闭
	defer rows.Close()
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Title, &book.Price)
		if err != nil {
			return
		}
		data = append(data, book)
	}
	return data
}

func insert(title string, price int) (err error) {
	// 更新 删除只用修改sql语句
	sqlStr := fmt.Sprintf(`insert into book (title,price) values ("%s",%d);`, title, price)
	ret, err := db.Exec(sqlStr)
	if err != nil {
		return
	}
	// 如果是插入操作，能拿到插入数据的id
	_, err = ret.LastInsertId()
	if err != nil {
		return
	}
	return
}

func delete(id int) (err error) {
	sqlStr := fmt.Sprintf(`delete from book where id=%d ;`, id)
	_, err = db.Exec(sqlStr)
	return
}
