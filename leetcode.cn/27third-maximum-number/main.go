package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	var a, b, c *int
	for _, num := range nums {
		num := num //改变内存地址
		if a == nil || num > *a {
			a, b, c = &num, a, b
		} else if num < *a && (b == nil || *b < num) {
			b, c = &num, b
		} else if b != nil && num < *b && (c == nil || *c < num) {
			c = &num
		}
	}
	if c != nil {
		return *c
	}
	return *a
}
func main() {
	fmt.Println(thirdMax([]int{5, 2, 2}))
}
