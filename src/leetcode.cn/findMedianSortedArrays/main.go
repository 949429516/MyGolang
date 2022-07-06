/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	mid := totalLength/2 + 1
	if totalLength%2 == 1 {
		// 奇数，两个切片长度为奇数则说明中位数是切片长度的中间值
		return float64(getK(nums1, nums2, mid))
	} else {
		// 偶数
		return float64(getK(nums1, nums2, mid)+getK(nums1, nums2, mid-1)) / 2.0
	}
}
func getK(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[k-1]
		}
		if index2 == len(nums2) {
			return nums1[k-1]
		}
		if k == 1 {
			return compareNumbers(nums1[index1], nums2[index2])
		}
		half := k / 2
		index1, index2 = half-1, half-1
		if index1 >= len(nums1) {
			index1 = len(nums1) - 1
		}
		if index2 >= len(nums2) {
			index2 = len(nums2) - 1
		}
		if nums1[index1] <= nums2[index2] {
			nums1 = nums1[index1+1:]
			k -= index1+1
			index1 = 0
		} else {
			nums2 = nums2[index2+1:]
			k -= index2+1
			index2 = 0
		}
	}
}
func compareNumbers(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	a := []int{1}
	b := []int{2, 3, 4, 5, 6}
	fmt.Println(findMedianSortedArrays(a, b))
}
