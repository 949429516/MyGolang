package splitstring

import (
	"strings"
)

// 根据sep切割字符串s
func SplitString(s, sep string) (result []string) {
	index := strings.Index(s, sep)
	sepLen := len(sep)
	for index > -1 {
		result = append(result, s[:index])
		s = s[index+sepLen:]
		index = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

// 优化内存申请
func SplitString1(s, sep string) []string {
	sepLen := len(sep)
	result := make([]string, 0, strings.Count(s, sep)+1)
	index := strings.Index(s, sep)
	for index > -1 {
		result = append(result, s[:index])
		s = s[index+sepLen:]
		index = strings.Index(s, sep)
	}
	result = append(result, s)
	return result
}

//斐波纳锲 验证性能比较函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
