package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxProfit(prices []int) int {
	// 如果切片长度小于等于1则收益为0
	if len(prices) <= 1 {
		return 0
	}
	// 创建二维切片
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(prices))
	}
	// 初始化dp，第一天的收益dp[0]和支出dp[1]
	dp[0][0] = 0
	dp[1][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[0][i] = max(dp[0][i-1], prices[i]+dp[1][i-1]) // 当前的股票价格prices[i]+昨天的支出==收益，与前一天的收益取最大值
		dp[1][i] = max(dp[1][i-1], -prices[i])           // 保存当前的支出
	}
	// 二维切片中的下标0切片最后一位就是最终的收益
	return dp[0][len(prices)-1]
}
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
