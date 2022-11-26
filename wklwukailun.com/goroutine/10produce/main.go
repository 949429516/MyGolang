package main

import (
	"fmt"
	"sync"
)

func doIt(workID int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running \n", workID)
	defer wg.Done()
	for {
		select {
		case m, ok := <-ch:
			if ok { //判断通道是否关闭
				fmt.Printf("[%v] m => %v\n", workID, m)
			}
		case <-done: // 只要接收到该通道发来的任何消息，及判定为收到关闭消息，该操作利用了通道关闭的广播消息
			fmt.Printf("[%v] is done\n", workID)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	ch := make(chan interface{}, 10)
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, ch, done, &wg) // wg 传递指针,doIt() 内部改变wg的值
	}
	// 主进程充当生产者角色向通道中发送消息
	for i := 0; i < workerCount; i++ { // 向ch中发送数据, 关闭goroutine
		ch <- i
	}
	// 关闭通道,生产者完成所有消息的生产之后主动关闭通道
	close(ch)
	// 关闭done通道，利用管波消息通知所有goroutine关闭信号
	close(done)
	wg.Wait()
	fmt.Println("all done")
}
