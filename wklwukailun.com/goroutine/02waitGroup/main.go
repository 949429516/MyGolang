package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup
func f() {
	rand.Seed(time.Now().UnixNano()) //种子,保证每次执行的数字不一样
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    //int64
		r2 := rand.Intn(10) //0<=x<10
		fmt.Println(r1, r2)
	}
}
func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	// 如何知道10个goroutine都结束了
	wg.Wait() //等待wg的计数器为0
}
