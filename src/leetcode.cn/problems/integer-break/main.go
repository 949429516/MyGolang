package main

/*
给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。

返回 你可以获得的最大乘积 。

示例 1:

输入: n = 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: n = 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
*/
import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curmax := 0
		for j := 1; j < i; j++ {
			curmax = max(curmax, max((i-j)*j, dp[i-j]*j))
		}
		dp[i] = curmax
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func main() {
	fmt.Println(integerBreak(4))
}
