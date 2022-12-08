package main

import "fmt"

func getNext(s string) []int {
	/*
		如果s[i] == s[j] 则指针j的位置+1 j+1
		如果j>0  且 s[i] != s[j] 则 指针j退回上一个位置的next数组的位置
		如果j==0 且 s[i] != s[j] 则 next数组i的位置就是j的值
	*/
	next := make([]int, len(s))
	next[0] = 0
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return -1
	}
	next := getNext(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}
func main() {
	// fmt.Println(strStr("aabaabaafa", "aabaaf"))
	fmt.Println(getNext("ababcaabc"))
}
