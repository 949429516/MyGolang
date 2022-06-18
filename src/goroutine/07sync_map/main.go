package main

/*
代码开启少量几个goroutine的时候可能没什么问题，
当并发多了之后执行上面的代码就会报
fatal error: concurrent map writes错误
*/
import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}
func set(key string, value int) {
	m[key] = value
}
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=%v,v=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
