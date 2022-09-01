package main

/*
以下代码有什么问题？如何解决？
fatal error: concurrent map writes 并发时会出现锁的问题
加锁

mu := &sync.RWMutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[rand.Int()] = rand.Int()
			mu.Unlock()

*/
import (
	"fmt"
	"math/rand"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			m[rand.Int()] = rand.Int()
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}
