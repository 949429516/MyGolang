package main

import "fmt"

func adder() func(int) int {
	var x int = 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder1(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder()
	rett := ret(200)
	fmt.Println(rett)

	ret1 := adder1(100)
	rett1 := ret1(200)
	fmt.Println(rett1)
}
