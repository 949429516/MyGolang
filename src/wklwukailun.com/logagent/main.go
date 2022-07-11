package main

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
	"wklwukailun.com/logagent/conf"
	"wklwukailun.com/logagent/kafka"
	"wklwukailun.com/logagent/taillog"
)

var (
	cfg = new(conf.AppConf)
)

// logAgent入口程序
func run() {
	//1 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2 发送kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}
func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed,err:", err)
		return
	}
	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.Address})
	if err != nil {
		fmt.Println("init kafka failed,err", err)
		return
	}
	fmt.Println("init kafka success!!!")
	// 2.打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Println("Init taillog failed,err:", err)
	}
	fmt.Println("init taillog success")
	run()
}

// kafka自带消费者 ./kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic=web_log --from-beginning
