package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	//修改的外层函数的base
	fmt.Println(f1(1), f2(2)) // 11  9
	fmt.Println(f1(1), f2(2)) // 10  8
}
