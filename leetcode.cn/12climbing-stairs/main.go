/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

动态规划：
1.上1个台阶的方法有1种
2.上2个台阶的方法有2种
3.上3个台阶的方法有3种
4.上4个台阶的方法有5种
.
.
.
每一个上台阶数都是由前两个推导出来的

定义dp
	dp[i]：爬到第i层楼梯，有dp[i]种方法。
	爬到第i层，可以从第i-1层上一个台阶；
	爬到第i层，可以从第i-2层上两个台阶；
	爬到第i层的总方法等于上面两种方法之和：dp[i] = dp[i - 1] + dp[i - 2]
*/

package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
func main() {
	fmt.Println(climbStairs(4))

}
