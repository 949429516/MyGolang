package main

import "fmt"

func assign(a interface{}) {
	fmt.Printf("类型是:%T\n", a)
	if str, ok := a.(string); ok {
		fmt.Println(str)
	} else {
		fmt.Println("不是字符型")
	}
}

func assignswitch(a interface{}) {
	switch t := a.(type) {
	default:
		fmt.Println("不在所属类型中")
	case int:
		fmt.Println(t)
	case string:
		fmt.Println(t)
	case bool:
		fmt.Println(t)
	}
}

func main() {
	assign(1)
	assign("你好")
	var a map[string]int
	assignswitch(a)
	assignswitch(123)
}
