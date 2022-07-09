package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// sarama第三方库开发kafka client
func main() {
	config := sarama.NewConfig()
	// tailf包使用
	// leader 和 follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完整数据需要,0 1 all三种模式
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个分区 partition 三种模式 指定key 轮循
	config.Producer.Return.Successes = true                   //成功交付消息将在success channel返回
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return
	}
	fmt.Printf("pid:%v,offset:%v\n", pid, offset)

}
