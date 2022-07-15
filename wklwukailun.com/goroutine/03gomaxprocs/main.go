package main

/*
可增长的栈

OS线程（操作系统线程）一般都有固定的栈内存(通常为2MB)，一个goroutine的栈在其生命周期开始时只有很小(2kb),
goroutine的栈不是固定的，他可以按需增加或减小，goroutine的栈的大小限制可以达到1GB，虽然很少会用到这个大小，
所以在Go语言中一次创建十万左右的goroutin也是可以。

goroutine调度
GMP是Go语言运行时(runtime)层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。
G：就是goroutine的，里面处了存放本goroutine信息外，还与坐在P的绑定等信息。

M：是go运行时（runtime）对操作系统内核线程的虚拟，M与内核线程一般是映射关系，一个goroutine最终是要放到M上执行的；

P：管理着一组goroutine队列，P里面会存储当前的goroutine运行的上下文环境（函数指针，堆栈地质及地址边界），P会对自己
管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行、后续的goroutine等）当自己的队列消费完毕
后就去全局队列中获取，如果全局队列里也消费完毕就会去其他P的队列里抢任务。
P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。 在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。

单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术
（复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着
一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上，
再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。

GMP
goroutine
p调度者
M:N把M个goroutine分配给n个操作系统线程去执行
goroutine初始栈的大小是2k
*/
/*
1.一个操作系统线程对应用户态多个goroutine
2.go程序可以同时使用多个操作系统线程
3.goroutine和OS线程是多对多的关系，即m:n
*/
import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	//runtime.GOMAXPROCS(4) //默认CPU的逻辑核心数
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
