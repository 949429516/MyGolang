package main

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。


示例 1：
输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
示例 2:
输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3：
输入：s = "ab", p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/regular-expression-matching
*/
import "fmt"

/*
if 是*号 {
	匹配0次(查看*号之前前的是否匹配成功) 或者  匹配多次
	dp[i][j] = dp[i][j-2]           ||   (dp[i-1][j] && (s[i-1]==p[j-2] || p[j-2]=='.'))
}else{
	不是*号,前一个s和p可以匹配的情况 && (当前的s和p相等 || 当前的p是.)
	dp[i][j] = dp[i-i][j-1] && (s[i-1]==p[j-1] || p[i-1]=='.')
}
*/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	// 处理dp的第一列*号
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(isMatch("aa", "a*"))
}
