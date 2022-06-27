package main

import "fmt"

func main() {
	//切片
	// var s1 []int //没有分配内存==nil
	// fmt.Println(s1)
	// fmt.Println(s1 == nil)
	// s1 = []int{1, 2, 3}
	// fmt.Println(s1)
	// //make初始化,分配内存
	// s2 := make([]bool, 2, 4)
	// fmt.Println(s2)
	// s3 := make([]int, 0, 4)
	// fmt.Println(s3, s3 == nil)

	//复制,赋值
	s1 := []int{1, 2, 3}
	s2 := s1
	var s3 = make([]int, 3, 3)
	copy(s3, s1)
	fmt.Println(s2)
	s2[1] = 200
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println(s3)
	var s4 []int
	s4 = append(s4, 1)
	fmt.Println(s4)
	var s5 = make([]int, 5, 5)
	s5 = []int{1, 2, 3, 4, 5}
	s6 := append(s5[0:1], s5[2:]...)
	fmt.Println(s6, s5)

	//指针,只能读取但不能修改
	name := "傻屄"
	nameP := &name
	fmt.Println(nameP)
	fmt.Printf("%T\n", nameP)
	nameV := *nameP
	fmt.Println(nameV)
	//map存储的是键值对
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["王森宝"] = 10
	fmt.Println(m1)
	fmt.Println(m1["wang"]) //如果获取不到key会返回对应类型的0
	socore, ok := m1["wang"]
	if !ok {
		fmt.Println("未获取")
	} else {
		fmt.Println(socore)
	}
	//删除,1.如果删除的不存在什么都不做
	delete(m1, "王森宝")
	fmt.Println(m1)
}
