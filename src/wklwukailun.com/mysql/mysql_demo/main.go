package main

//go get -u github.com/go-sql-driver/mysql

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

// 查询一条
func queryRowOne(id int) {
	var u user
	sqlStr := `select id,name,age from user where id=?;`
	//执行
	rowObj := db.QueryRow(sqlStr, id) //从连接池里拿出连接去查询数据库
	//拿到结果
	rowObj.Scan(&u.id, &u.name, &u.age) //取保QueryRow之后调用Scan方法,否则持有数据库连接不会释放会阻塞
	fmt.Println(u)
}

// 查询多条
func queryRowMore(id int) {
	sqlStr := `select id,name,age from user where id>?;`
	rows, err := db.Query(sqlStr, id)
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

// 插入数据
func insert() {
	// 更新 删除只用修改sql语句
	sqlStr := `insert into user(name,age) values ("王小雅",20);`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		return
	}
	// 如果是插入操作，能拿到插入数据的id
	id, err := ret.LastInsertId()
	//n,err :=ret.RowsAffected() //更新多少数据
	if err != nil {
		return
	}
	fmt.Println("id:", id)
}

// 预处理方式插入多条数据
func prepareInsert() {
	/*
			普通SQL语句执行过程：

		客户端对SQL语句进行占位符替换得到完整的SQL语句。
		客户端发送完整SQL语句到MySQL服务端
		MySQL服务端执行完整的SQL语句并将结果返回给客户端。
		预处理执行过程：

		把SQL语句分成两部分，命令部分与数据部分。
		先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
		然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
		MySQL服务端执行完整的SQL语句并将结果返回给客户端。
		为什么要预处理？
		优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
		避免SQL注入问题。
	*/
	sqlStr := `insert into user(name,age) values(?,?);`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer stmt.Close()
	var m = map[string]int{
		"赵恒":  67,
		"王楚楚": 29,
	}
	// 后续只用stmt来操作
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

//连接mysql
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	//queryRowOne(1)
	//insert()
	prepareInsert()
	queryRowMore(0)
}
