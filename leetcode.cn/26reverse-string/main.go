package main

import "fmt"

func reverseString(s []string) []string {

	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
	return s
}

func main() {
	fmt.Println(reverseString([]string{"h", "e", "l", "l", "o"}))
	fmt.Println(reverseString([]string{"s", "v", "c"}))
	fmt.Println(reverseString([]string{"H", "a", "n", "n", "a", "h"}))
}
