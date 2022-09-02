package main

/*
下面代码输出什么，如何让输出为true
输出是false;改为值类型
*/

import "fmt"

type S struct {
	a, b, c string
}

func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})
	fmt.Println(x == y)
}
