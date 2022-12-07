package main

import "fmt"

func getNext(s string) []int {
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

func getNext1(s string) []int {
	next := make([]int, len(s))
	next[0] = 0
	for i, j := 1, 0; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j--
		}
		if s[i] == s[j] {
			next[i] = next[j] + 1
			j++
		}
		if j == 0 && s[i] != s[j] {
			next[i] = 0
		}
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
	fmt.Println(getNext1("ababcaabc"))
}
