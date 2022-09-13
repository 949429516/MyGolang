package main

/*
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :
输入: 2736
输出: 7236
解释: 交换数字2和数字7。

示例 2 :
输入: 9973
输出: 9973
解释: 不需要交换。
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-swap

生成两个切片，原顺序original，从大到小排序sorted
[1 9 9 5]
[9 9 5 1]
顺序比较两个切片，如果出现不相等则这就是要交换的
将原切片的数字和对应的sorted切片的数字对应原切片数字的位置交换（比如1和十位的9交换）
*/
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func maximumSwap(num int) (ret int) {
	original := []string{}
	for _, i := range strconv.Itoa(num) {
		original = append(original, string(i))
	}
	sorted := make([]string, len(original))
	copy(sorted, original)
	sort.Sort(sort.Reverse(sort.StringSlice(sorted)))
	for i := 0; i < len(original); i++ {
		if original[i] != sorted[i] {
			index := 0
			for index_i, value := range original {
				if value == sorted[i] {
					index = index_i
				}
			}
			original[i], original[index] = original[index], original[i]
			break
		}
	}
	s := strings.Join(original, "")
	ret, _ = strconv.Atoi(s)
	return
}

func main() {
	fmt.Println(maximumSwap(1995))
}
