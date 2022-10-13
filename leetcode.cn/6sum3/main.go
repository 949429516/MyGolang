package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) (ret [][]int) {
	if len(nums) < 3 {
		return
	}
	sort.Sort(sort.IntSlice(nums))
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		right := len(nums) - 1
		left := i + 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else if sum == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 由于是从小到大排序，当前已经为0时左右需要同时改变
				left++
				right--
			}
		}
	}
	return
}
func main() {
	// [-1,0,1,2,-1,-4,-2,-3,3,0,4]
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
