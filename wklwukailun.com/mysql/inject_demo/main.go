package main

// sql注入
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init
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
	//defer db.Close() // 注意这行代码要写在上面err判断的下面
}
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id,name,age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	rows, err := db.Query(sqlStr)
	if err != nil {
		return
	}
	//一定要关闭
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			return
		}
		fmt.Println(u)
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	// sql注入的几种案例
	sqlInjectDemo("王森堡")
	sqlInjectDemo("xxx' or 1=1#")
	sqlInjectDemo("xxx' union select * from user #")

}
