/*
编写一个算法来判断一个数 n 是不是快乐数。
「快乐数」 定义为：
对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false
*/
package main

import "fmt"

func isHappy(n int) bool {
	//设置两个快慢
	slow, fast := n, n
	for {
		// 快慢任何一个先到1则说明是快乐数
		if slow == 1 || fast == 1 {
			return true
		}
		// 慢的每次走一步，快的每次走两步
		slow = getNext(slow)
		fast = getNext(getNext(fast))
		// 如果两个相等则说明出现循环了
		if slow == fast && slow != 1 {
			return false
		}
	}
}
func getNext(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
func main() {
	fmt.Println(isHappy(20))
}
