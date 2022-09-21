package main

/*
我们给出了一个（轴对齐的）二维矩形列表 rectangles 。 对于 rectangle[i] = [x1, y1, x2, y2]，
其中（x1，y1）是矩形 i 左下角的坐标， (xi1, yi1) 是该矩形 左下角 的坐标， (xi2, yi2) 是该矩形 右上角 的坐标。
计算平面中所有 rectangles 所覆盖的 总面积 。任何被两个或多个矩形覆盖的区域应只计算 一次 。
返回 总面积 。因为答案可能太大，返回 10的9次方 + 7 的 模 。

示例 1：
输入：rectangles = [[0,0,2,2],[1,0,2,3],[1,0,3,1]]
输出：6
解释：如图所示，三个矩形覆盖了总面积为6的区域。
从(1,1)到(2,2)，绿色矩形和红色矩形重叠。
从(1,0)到(2,3)，三个矩形都重叠。
示例 2：

输入：rectangles = [[0,0,1000000000,1000000000]]
输出：49
解释：答案是 1018 对 (109 + 7) 取模的结果， 即 49 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/rectangle-area-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import (
	"fmt"
	"sort"
)

func rectangleArea(rectangles [][]int) (ret int) {
	// 1.将x轴的所有值取出存入切片并且从小到大排序
	x_slice := []int{}
	for _, slice := range rectangles {
		x_slice = append(x_slice, slice[0], slice[2])
	}
	sort.Sort(sort.IntSlice(x_slice))
	// 2.x轴切割出的每一个“宽度”;
	for i := 0; i < len(x_slice)-1; i++ {
		a, b := x_slice[i], x_slice[i+1]
		width := b - a
		if width == 0 { // 宽度为0则不用计算
			continue
		}
		// 2.1遍历切片找到与该“宽度”相交的矩形，得到y轴的低点和高点,存入二维切片
		lines := [][]int{}
		for _, x := range rectangles {
			if a >= x[0] && b <= x[2] {
				lines = append(lines, []int{x[1], x[3]})
			}
		}
		// 2.2从小到大排序
		sort.Slice(lines, func(i, j int) bool {
			if lines[i][0] == lines[j][0] {
				return lines[i][1] <= lines[j][1]
			}
			return lines[i][0] <= lines[j][0]
		})
		// 3.定义当前循环的高度hight,初始y轴low和high,遍历高度二维切片
		hight, low, high := 0, 0, 0
		for _, current := range lines {
			// 4.如果high<=当前y轴最低点(说明矩形中有空白部分不计算)，计算出存储的high-low高度并且增加到hight;更新low和high(当前遍历到切片的最低点和最高点)
			if high < current[0] {
				hight += high - low
				low, high = current[0], current[1]
				// 5.如果high<=当前y轴最高点(说明当前的高度包含了之前的高度)，更新high为当前最高
			} else if high < current[1] {
				high = current[1]
			}
		}
		// 最后一次循环没有增加的在循环退出的时候计算
		hight += high - low
		ret += width * hight
	}
	ret = ret % 1000000007
	return
}

func main() {
	fmt.Println(rectangleArea([][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}))
}
