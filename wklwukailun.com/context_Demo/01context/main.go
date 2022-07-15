package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("王森堡")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break LOOP
		default:

		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	exitChan <- true
	wg.Wait()
	//如何通知子goroutine退出
}
