package main

/*
互斥锁(防止同一时刻多个goroutine操作同一个资源)
是完全互斥的，很多场景下读多写少，当我们读取一个资源不涉及写时是没必要加锁的,读写时会影响性能
常用的控制共享资源的方法，保证只有一个gorotine访问共享资源
*/
import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
