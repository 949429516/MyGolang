package main

import (
	"fmt"
	"sync"
)

/*
代码输出是什么？为什么？如何修改使len(m)为10
1
原因是for循环速度很快，循环内部的协程获取的是外部作用域的i,协程启动的时候可能获取的都是同一个值
加上map的key不可重复所以出现覆盖从而长度为1

for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
*/

const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	// for循环的问题导致，
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}
