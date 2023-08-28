package main

import "fmt"

func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		x := (columnNumber-1)%26 + 1
		ans = append(ans, 'A'+byte(x-1))
		columnNumber = (columnNumber - x) / 26
	}
	for i, j := 0, len(ans)-1; i < len(ans)/2; i++ {
		ans[i], ans[j-i] = ans[j-i], ans[i]
	}
	return string(ans)
}
func main() {
	fmt.Println(convertToTitle(701))
}
