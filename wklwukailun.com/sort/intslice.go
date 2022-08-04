package main

import (
	"fmt"
	"sort"
)

func main() {
	demo := []int{10, 1, 2, 9, 8, 3, 4, 7, 5, 6}
	fmt.Println("排序前:", demo)
	ok := sort.IsSorted(sort.IntSlice(demo))
	fmt.Println("是否已经排序", ok)
	sort.Sort(sort.IntSlice(demo))
	fmt.Println("排序后:", demo)
	ok = sort.IsSorted(sort.IntSlice(demo))
	fmt.Println("是否已经排序", ok)
	key := sort.SearchInts(demo, 2)
	fmt.Println("查询al为2的key是:", key)
	sort.Sort(sort.Reverse(sort.IntSlice(demo)))
	fmt.Println("反转后:", demo)
}
