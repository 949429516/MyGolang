package main

import "fmt"

func main() {

	m1 := make(map[string]interface{}, 16)
	m1["姓名"] = "李沁岩"
	m1["age"] = 30
	m1["merried"] = false
	m1["hobby"] = [...]string{"约炮", "嫖娼", "蹦迪"}
	fmt.Println(m1)
}
