package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	//当子goroutine又开启另外一个goroutine时，只需要将ctx传入即可：
	go f2(ctx)
LOOP:
	for {
		fmt.Println("李沁演")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}
func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("王森堡")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	//如何通知子goroutine退出
}
