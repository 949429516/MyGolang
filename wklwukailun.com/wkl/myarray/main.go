package main

import "fmt"

func main() {
	var x = [3]int{1, 2, 3}
	fmt.Println(x)
	a := f1(x)
	fmt.Println(a)
}

func f1(a [3]int) [3]int {
	a[1] = 100
	return a
}
