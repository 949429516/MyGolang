package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	//必须首字母大写，不大写json获取不到
	//小写字母开头的只能在本main包里获取到，json模块在别的包里
	Name string `json:"name"` //json返回时命名修改
	Age  int
}

func main() {
	p1 := person{
		Name: "王傻屄",
		Age:  38,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err%v", err)
	} else {
		fmt.Println(string(b))
	}
	//反序列化
	str := `{"name":"王傻屄","Age":38}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了在函数内部修改
	fmt.Printf("%#v", p2)
}
