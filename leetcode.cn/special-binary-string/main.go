/*
特殊的二进制序列是具有以下两个性质的二进制序列：
0 的数量与 1 的数量相等。
二进制序列的每一个前缀码中 1 的数量要大于等于 0 的数量。
给定一个特殊的二进制序列 S，以字符串形式表示。定义一个操作 为首先选择 S 的两个连续且非空的特殊的子串，然后将它们交换。（两个子串为连续的当且仅当第一个子串的最后一个字符恰好为第二个子串的第一个字符的前一个字符。)
在任意次数的操作之后，交换后的字符串按照字典序排列的最大的结果是什么？

示例 1:
输入: S = "11011000"
输出: "11100100"
解释:
将子串 "10" （在S[1]出现） 和 "1100" （在S[3]出现）进行交换。
这是在进行若干次操作后按字典序排列最大的结果。
=========================================说人话==================================================
我们可以把特殊的二进制序列看作"有效的括号"，1代表左括号，0代表右括号。
0的数量与1的数量相等，代表左右括号数量相等。
二进制序列的每一个前缀码中1的数量要大于等于0的数量，代表有效的括号，每一个左括号都有右括号匹配，并且左括号在前。
比如："11011000"可以看作"(()(()))"。

两个连续且非空的特殊的子串，然后将它们交换，代表着交换两个相邻的两个有效括号。

我们可以将题进行如下划分，把每一个有效的括号匹配都看作一部分，然后进行排序，内部也进行排序处理，例如：

链接：https://leetcode.cn/problems/special-binary-string/
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {
	// 1.如果切割出的字符串<=2则返回
	if len(s) <= 2 {
		return s
	}
	// 2.创建一个排序切片字符串,将切割出的字符串从达到小排序
	slice := sort.StringSlice{}
	left, cur := 0, 0 // left为循环的左侧位置，cur为计数当为0时进入递归
	for index, b := range s {
		// 3. '1'相当于左括号+1，'0'相当于右括号-1
		if b == '1' {
			cur++
		} else {
			cur--
		}
		// 4.当cur为0时候则说明找到了闭合，将其闭合“括号”内的继续递归计算，将结果加入slice中
		if cur == 0 {
			slice = append(slice, "1"+makeLargestSpecial(s[left+1:index])+"0")
			left = index + 1
		}
	}
	sort.Sort(sort.Reverse(slice))
	return strings.Join(slice, "")
}

func main() {
	fmt.Println(makeLargestSpecial("11011000"))
}
