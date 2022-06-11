package main

import "fmt"

func f1() {
	fmt.Println("草泥马王森宝")
}

func f2(name string) {
	fmt.Println("hello", name)
}
func f3(x, y int) int {
	return x + y
}
func f4(x string, y ...int) int {
	fmt.Println(y) //y是一个int类型的切片可以写多个值
	return 1
}
func f5(x, y int) (sum int) {
	sum = x + y //如果使用命名的返回值，那么可以在函数中直接使用返回
	return      //可以省略返回值变量
}
func f6(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}
func main() {
	for i := 0; i < 10; i++ {
		f1()
	}
	fmt.Println(f3(100, 200))

	f4("贼尼玛", 1, 2, 3, 4, 5, 6)

}
