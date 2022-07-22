package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type logData struct {
	topic string
	data  string
}

// 专门向kafka写日志的模块
var (
	// 声明一个全局连接kafka的生产者client
	client sarama.SyncProducer
	// 通道保存发送过来的日志
	logDataChan chan *logData
)

// Init初始化client
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	// tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息在success channel返回

	//连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	// 初始化logDataChan
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine从通道中获取数据发往kafka
	go sendToKafka()
	return
}

// 给外部暴露的一个函数，该函数值只把日志发送到一个内部的channel中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

// 实际向kafka发送消息的函数
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			//构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			//发送到kafka
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				return
			}
			fmt.Printf("pid:%v,offset:%v,msg:%v\n", pid, offset, ld.data)
		default:
			time.Sleep(time.Millisecond * 50)
		}

	}
}

func SendToKafka(topic, data string) {
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	//发送到kafka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		return
	}
	fmt.Printf("pid:%v,offset:%v\n", pid, offset)
}
