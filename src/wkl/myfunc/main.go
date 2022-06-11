package main

import "fmt"

func sum(x int, y int) (ret int) {
	return x + y
}

func f(x int, y int) (ret int) {
	ret = x + y
	return
}

func ff(x int, y int) int {
	ret := x + y
	return ret
}

func f2(x int, y int) {
	fmt.Println(x + y)
}

func f3() int {
	return 3
}

func f4() (int, int) {
	return 1, 2
}

func f5(x, y int, m, n string) int {
	return 1
}

func f6(x string, y ...int) int {
	return 1
}
func main() {
	var r = sum(1, 2)
	fmt.Println(r)

	m, n := f4()
	fmt.Println(m, n)
}
