package main

import (
	"fmt"
)

/*
https://leetcode.cn/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：

输入：n = 1
输出：["()"]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func generateParenthesis(n int) (ret []string) {
	var f func(s string, left, right int)
	str := ""
	f = func(s string, left, right int) {
		if left == 0 && right == 0 {
			ret = append(ret, s)
			return
		}
		// 左侧括号数量大于右侧括号数量，则排除。因为，括号是字符串追加左侧多于右侧则说明后面追加的永远不会闭合)()()()(
		if left > right {
			return
		}
		if left > 0 {
			f(s+"(", left-1, right)
		}
		if right > 0 {
			f(s+")", left, right-1)
		}
	}
	f(str, n, n)
	return
}

func main() {
	fmt.Println(generateParenthesis(2))
}
