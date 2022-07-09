package main

import (
	"fmt"
	"time"

	"wklwukailun.com/logagent/kafka"
	"wklwukailun.com/logagent/taillog"
)

// logAgent入口程序
func run() {
	//1 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2 发送kafka
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}
func main() {
	// 1.初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Println("init kafka failed,err", err)
		return
	}
	fmt.Println("init kafka success!!!")
	// 2.打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Println("Init taillog failed,err:", err)
	}
	fmt.Println("init taillog success")
	run()
}

// kafka自带消费者 ./kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic=web_log --from-beginning
