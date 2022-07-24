package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"wklwukailun.com/logtransfer/es"
)

// Init初始化client，从kafka取数据发送ES
func Init(addrs []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		// defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 发送ES
				// ld := map[string]interface{}{
				// 	"data": string(msg.Value),
				// }
				ld := es.LogData{Topic: topic, Data: string(msg.Key)}
				// 优化:放到一个chan中
				es.SendToESChan(&ld)
			}
		}(pc)
	}
	return err
}
