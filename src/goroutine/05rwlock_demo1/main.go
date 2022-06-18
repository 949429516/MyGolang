package main

import (
	"fmt"
	"sync"
	"time"
)

/*
读写互斥锁
当一个goroutine获取读锁后其他的goroutine如果是获取读锁会继续获得锁,如果是获取写锁就会等待;
当一个goroutine读写锁之后,其他的goroutine无论是获取读锁还是写锁都会等待.
1.读的goroutine来访问的是读锁，后续的goroutine能读不能写
2.写的goroutine来访问的是写锁，后续的goroutine无论读写均要等待
*/
var (
	x      = 0
	wg     sync.WaitGroup
	rwlock sync.RWMutex //值类型，给函数传参时要传递指针类型。否则就是相当于复制(当前案例用全局变量可以不用指针)
)

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}
func write() {
	defer wg.Done()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
