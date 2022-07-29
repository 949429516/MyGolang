/*
给定一个表示分数加减运算的字符串 expression ，你需要返回一个字符串形式的计算结果。
这个结果应该是不可约分的分数，即最简分数。 如果最终结果是一个整数，例如 2，你需要将它转换成分数形式，其分母为 1。所以在上述例子中, 2 应该被转换为 2/1。

示例 1:
输入: expression = "-1/2+1/2"
输出: "0/1"
示例 2:
输入: expression = "-1/2+1/2+1/3"
输出: "1/3"
示例 3:
输入: expression = "1/3-1/2"
输出: "-1/6"
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/fraction-addition-and-subtraction
*/
package main

import (
	"fmt"
	"unicode"
)

// 分数计算 x1/y1 + x2/y2 = x1*y2+x2*y1 / y1*y2
func fractionAddition(expression string) string {
	// 0.定义分子 分母；遍历得到完整的一个分数后与fz fm计算出值并更新，循环完毕后fz fm就是要的值
	fz, fm := 0, 1
	for index, n := 0, len(expression); index < n; {
		// 1.获取分子
		// 1.1 分子的正负获取
		fz1, sign := 0, 1
		if expression[index] == '+' || expression[index] == '-' {
			if expression[index] == '-' {
				sign = -1
			}
			index++
		}
		// 1.2 获取分子数值
		// 由于数字中有10,10不是一个字节所以需要用for循环取数字
		for index < n && unicode.IsDigit(rune(expression[index])) {
			// *10是为了保证遇到10的时候特殊处理
			fz1 = fz1*10 + int(expression[index]-'0') // byte减去字节'0'转换int就是要的值或者使用strconv转换
			index++
		}
		// 1.3 得到分子，并且跨过/号
		fz1 = sign * fz1
		index++
		// 2.获取分母
		fm1 := 0
		for index < n && unicode.IsDigit(rune(expression[index])) {
			fm1 = fm1*10 + int(expression[index]-'0')
			index++
		}
		// 3.计算值并更新fz fm
		fz = fz*fm1 + fz1*fm
		fm = fm * fm1
	}
	if fz == 0 {
		return "0/1"
	}
	b := gcd(abs(fz), fm)
	return fmt.Sprintf("%d/%d", fz/b, fm/b)
}

// 绝对值
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// 最大公约数
func gcd(x, y int) int {
	for x != 0 {
		x, y = y%x, x
	}
	return y
}
func main() {
	fmt.Println(fractionAddition("1/2+1/3"))
}
