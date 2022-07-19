/*
实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。
MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。
当三个日程安排有一些时间上的交叉时（例如三个日程安排都在同一时间内），就会产生三重预订。
每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致三重预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。
请按照以下步骤调用MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)
示例：
MyCalendar();
MyCalendar.book(10, 20); // returns true
MyCalendar.book(50, 60); // returns true
MyCalendar.book(10, 40); // returns true
MyCalendar.book(5, 15); // returns false
MyCalendar.book(5, 10); // returns true
MyCalendar.book(25, 55); // returns true
解释：
前两个日程安排可以添加至日历中。 第三个日程安排会导致双重预订，但可以添加至日历中。
第四个日程安排活动（5,15）不能添加至日历中，因为它会导致三重预订。
第五个日程安排（5,10）可以添加至日历中，因为它未使用已经双重预订的时间10。
第六个日程安排（25,55）可以添加至日历中，因为时间 [25,40] 将和第三个日程安排双重预订；
时间 [40,50] 将单独预订，时间 [50,55）将和第二个日程安排双重预订。
*/
package main

import "fmt"

// 存放数轴位置
type pair struct {
	start int
	end   int
}
type MyCalendarTwo struct {
	// 存放没有重合的数轴
	books []pair
	// 存二重合的数轴
	double []pair
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}
func (this *MyCalendarTwo) Book(start, end int) bool {
	for _, i := range this.double {
		// 判断新加入的数轴有没有与所有二重和的数轴重合，若重合则为三重和返回false
		if start < i.end && i.start < end {
			// [start,end) [i.start,i.end) 两个数轴,目的是有重合的返回false;因为[i.start,i.end)已经是二重和所以再重合就返回
			return false
		}
	}
	for _, i := range this.books {
		// 如果上面的循环遍历结束没有重合的,就在books中找二重和加入进去
		if start < i.end && i.start < end {
			// [start,end) [i.start,i.end) 两个数轴,目的是有重合的返回false;
			this.double = append(this.double, pair{max(start, i.start), min(end, i.end)})
		}
	}
	// 若上面两个循环都没有则直接加入books中,并返回
	this.books = append(this.books, pair{start, end})
	return true
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	MyCalendar := Constructor()
	fmt.Println(MyCalendar.Book(10, 20))
	fmt.Println(MyCalendar.Book(50, 60))
	fmt.Println(MyCalendar.Book(10, 40))
	fmt.Println(MyCalendar.Book(5, 15))
	fmt.Println(MyCalendar.Book(5, 10))
	fmt.Println(MyCalendar.Book(25, 55))
}
