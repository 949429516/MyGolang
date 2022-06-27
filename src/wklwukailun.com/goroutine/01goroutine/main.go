package main

/*
并发：同一时间段内执行多个任务（你在用微信和两个女朋友聊天）。

并行：同一时刻执行多个任务（你和你朋友都在用微信和女朋友聊天）。

Go语言的并发通过goroutine实现。goroutine类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个goroutine并发工作。
goroutine是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。

Go语言还提供channel在多个goroutine间进行通信。goroutine和channel是 Go 语言秉承的 CSP（Communicating Sequential Process）并发模式的重要实现基础。


在java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，
这一切通常会耗费程序员大量的心智。那么能不能有一种机制，程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？

Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine
中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，
开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。
*/

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

//程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 1000; i++ {
		//go hello(i) //开启一个单独的goroutine去执行hello函数(任务)
		go func(i int) {
			fmt.Println(i)
		}(i) //用的是该匿名函数自己的i,不是外层的i
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	//main函数结束后，由main函数启动的goroutine也都结束了
}
