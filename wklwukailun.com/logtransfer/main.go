package main

import (
	"fmt"

	"gopkg.in/ini.v1"
	"wklwukailun.com/logtransfer/conf"
	"wklwukailun.com/logtransfer/es"
	"wklwukailun.com/logtransfer/kafka"
)

// 将日志数据从kafka获取后发往ES

func main() {
	// 0.加载配置文件
	var cfg conf.LogTransfer
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Println("load ini failed,err:", err)
		return
	}
	// 2.初始化ES
	// 2.1初始化一个ES连接client
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.ChanNums)
	if err != nil {
		fmt.Println("init es failed,err", err)
		return
	}
	fmt.Println("init es success.")
	// 1.初始化kafka发送ES
	// 1.1连接kafka创建分区消费者
	// 1.2每个分区的消费者分别取出数据，将数据通过SendToES发往ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Println("init kafka consumer failed,err", err)
		return
	}
	select {}
}
