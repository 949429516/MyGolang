package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	//将defer后面的语句延迟到函数即将返回后再执行
	//例如在关闭文件的时候使用避免内存泄漏(释放资源)
	//如果有多个defer语句则遵守“先进后出”原则
	defer fmt.Println("嗨嗨嗨")
	defer fmt.Println("呵呵呵")
	defer fmt.Println("哈哈哈")
	fmt.Println("end")
}

func main() {
	deferDemo()
}
