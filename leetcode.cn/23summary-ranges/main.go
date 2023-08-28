package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) (res []string) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if fast+1 < len(nums) && nums[fast]+1 == nums[fast+1] {
			fast++
			continue
		}
		if fast == slow {
			res = append(res, strconv.Itoa(nums[fast]))
			fast++
			slow++
		} else {
			res = append(res, strconv.Itoa(nums[slow])+"->"+strconv.Itoa(nums[fast]))
			fast++
			slow = fast
		}
	}
	return
}

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	for {
		if n == 1 {
			return true
		} else if n%2 != 0 {
			return false
		}
		n = n / 2
	}
}
func main() {
	a := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(a))
	fmt.Println(isPowerOfTwo(-16))
}
