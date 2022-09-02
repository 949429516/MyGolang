package main

import (
	"fmt"
)

/*
下面代码的问题是什么，要如何修改？
type S struct {
	name string
}

func main() {
	m := map[string]S{"x": S{"one"}}
	m["x"].name = "two"
}

这是因为map 中的元素是不可寻址的。由于map 会进行动态扩容，当进行扩展后，
map的value就会进行内存迁移，其地址发生变化，因此没法对这个value进行寻址
*/
type S struct {
	name string
}

func main() {
	//定义map数据的时候使用结构体指针，
	//因为存的是结构体的地址所以这个时候我们自然也就能修改结构体的值了。
	m := map[string]*S{"x": &S{name: "one"}}
	m["x"].name = "two"
	fmt.Println(m["x"].name)
	// 利用局部变量来完成修改
	m1 := map[string]S{"x": S{name: "one"}}
	r := m1["x"]
	r.name = "two"
	m1["x"] = r
	fmt.Println(m1["x"].name)
}
