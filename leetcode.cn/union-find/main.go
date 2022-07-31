package main

/*
给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：
有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
返回 图中最大连通组件的大小 。
示例 1：
输入：nums = [4,6,15,35]
输出：4
示例 2：
输入：nums = [20,50,9,63]
输出：2
示例 3：
输入：nums = [2,3,6,7,4,12,21,39]
输出：8
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/largest-component-size-by-common-factor
*/
import "fmt"

// 合并集

type unionFind struct {
	father []int // 数组存储，下标作为实际值，值作为其指向的父亲节点
	rank   []int // 按秩合并，含义是较少节点的指向较多节点，这样可以优化路径合并的复杂度
}

// 根据切片中最大数字，初始化结构体生成合并集结构
func newUnionFind(n int) *unionFind {
	father := make([]int, n)
	for index := range father {
		// 遍历下标
		father[index] = index
	}
	return &unionFind{father: father, rank: make([]int, n)}
}

// 查找父亲节点
func (union *unionFind) find(x int) int {
	// 递归函数,如果index==x则说明找到父亲节点
	if union.father[x] != x {
		union.father[x] = union.find(union.father[x]) // 路径压缩
	}
	return union.father[x]
}

// 合并
func (union *unionFind) mearge(x, y int) {
	x, y = union.find(x), union.find(y)
	if x == y {
		// 若两个值的父节点相等则说明在同一个集合中，直接返回
		return
	}
	// 按秩合并
	if union.rank[x] > union.rank[y] {
		// 若x集合大于y集合,让y集合指向x集合,y的父亲节点就是x
		union.father[y] = x
		union.rank[x]++
	} else {
		union.father[x] = y
		union.rank[y]++
	}

}

// 找出最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func largestComponentSize(nums []int) (ret int) {
	maxDigit := 0
	// 遍历切片的值,找出最大值作为并查集的初始化结构
	for _, i := range nums {
		maxDigit = max(i, maxDigit)
	}
	un := newUnionFind(maxDigit + 1) // 遍历下标是从0开始则+1
	// 找到nums中每个值的因数，加入并查集中
	for _, num := range nums {
		// 1不是因数,且num/i也是因数,因数范围在[2,根num]
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				un.mearge(num, i)
				un.mearge(num, num/i)
			}
		}
	}
	// 并查集生成完毕后，找到最大的集合并且计数(找到包含元素最多父节点)
	ret_slice := make([]int, maxDigit+1)
	for _, num := range nums {
		fa := un.find(num)            // 找其的父节点
		ret_slice[fa]++               // 给num的父节点+1
		ret = max(ret, ret_slice[fa]) // 当前最大的值赋值给ret最终遍历完成后返回的就是最大
	}
	return
}

func main() {
	fmt.Println(largestComponentSize([]int{4, 6, 15, 35}))
}
