package main

import "fmt"

/*闭包
闭包是一个函数，这个函数包含了外部作用域的一个变量
底层原理
1.函数可以作为返回值
2.函数内部查找变量的顺序，先在内部查找，找不到去外层找
*/
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2包装
func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}

func main() {
	ret := f3(f2, 1, 2)
	ret()
	f1(ret)
}
