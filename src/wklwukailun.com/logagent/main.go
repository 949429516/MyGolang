package logagent

import (
	"fmt"

	"wklwukailun.com/logagent/kafka"
	"wklwukailun.com/logagent/taillog"
)

// logAgent入口程序
func run() {
	//1 读取日志
	//2 发送kafka
}
func main() {
	// 1.初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Println("init kafka failed,err", err)
		return
	}
	// 2.打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Println("Init taillog failed,err:", err)
	}
	run()
}
