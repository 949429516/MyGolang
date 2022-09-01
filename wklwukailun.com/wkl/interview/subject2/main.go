package main

import "fmt"

/*
下面代码输出是什么，想输出012，怎么改
指针类型是同一个i会一直覆盖
3
3
3
不传指针就可以
*/

const N = 3

func main() {
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		m[i] = &i
	}
	for _, v := range m {
		fmt.Println(*v)
	}
}
