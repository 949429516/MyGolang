/*
二分法
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sqrtx
*/
package main

import "fmt"

func mySqrt(x int) (ans int) {
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid > x {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return
}

func main() {
	fmt.Println(mySqrt(9))
}
