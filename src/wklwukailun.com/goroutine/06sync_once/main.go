package main

import (
	"fmt"
	"sync"
)

/*
编程中很多场景需要确保只执行一次
*/
var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	f := func() {
		close(ch2)
	}
	once.Do(f) //只执行一次,没有参数的函数
}
func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
