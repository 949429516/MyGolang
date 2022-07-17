package main

import (
	"fmt"
	"sync"
	"time"

	"gopkg.in/ini.v1"
	"wklwukailun.com/logagent/conf"
	"wklwukailun.com/logagent/etcd"
	"wklwukailun.com/logagent/kafka"
	"wklwukailun.com/logagent/taillog"
)

var (
	cfg = new(conf.AppConf)
)

// logAgent入口程序
func run() {
	//1 读取日志
	// for {
	// 	select {
	// 	case line := <-taillog.ReadChan():
	// 		//2 发送kafka
	// 		kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
	// 	default:
	// 		time.Sleep(time.Second)
	// 	}
	// }

}
func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed,err:", err)
		return
	}
	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Println("init kafka failed,err", err)
		return
	}
	fmt.Println("init kafka success!!!")
	// 2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("connect to etcd failed, err:", err)
		return
	}
	fmt.Println("connect to etcd success")
	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Println("get conf failed, err:", err)
		return
	}
	fmt.Println("get conf from etcd success", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Println(index, value)
	}
	// 2.打开日志文件准备收集日志
	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Println("Init taillog failed,err:", err)
	// }
	// fmt.Println("init taillog success")
	// run()
	// 3.收集日志发往kafak
	// 3.1 循环每一个收集项，创建tailobj 3.2 发往kafka
	taillog.Init(logEntryConf)
	// 3.2 派哨兵监视日志收集项的变化(有变化及时通知为的logAgent实现日志加载配置)
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan) // 哨兵发现最新的配置信息通知上面的通道
	wg.Wait()
}

// kafka自带消费者 ./kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic=web_log --from-beginning
