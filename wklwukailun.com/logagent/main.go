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
	"wklwukailun.com/logagent/utlis"
)

var (
	cfg = new(conf.AppConf)
)

// logAgent入口程序
func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed,err:", err)
		return
	}
	// 1.初始化kafka连接;1.1 开启一个logDataChan通道存放topic和data;1.2开启goroutine从logDataChan通道中获取数据发往kafka
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
	// 2.1 从etcd中通过key获取日志收集项的配置信息,将配置信息转换为json格式存入切片结构体logEntry,返回logEntry切片(存放etcd发现的配置信息)
	// 为了实现每个logagnet都拉取自己独有的配置，需要以自己的ip地址作为区分
	ipStr, err := utlis.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Println("get conf failed, err:", err)
		return
	}
	fmt.Println("get conf from etcd success", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Println(index, value)
	}
	// 3.收集日志发往kafak
	// 3.初始化taillog;3.1.创建一个taillog的管理者;3.2.遍历logEntry;3.3.打开文件读取;3.4发送kafka的logDataChan通道中
	taillog.Init(logEntryConf)
	// 3.2 派哨兵监视日志收集项的变化(有变化及时通知为的logAgent实现日志加载配置)
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) // 哨兵发现最新的配置信息通知上面的通道
	wg.Wait()
}

// kafka自带消费者 ./kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic= --from-beginning
