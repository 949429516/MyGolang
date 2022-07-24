package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	// 全局ES客户端
	client *elastic.Client
	ch     chan *LogData
)

// 初始化ES,准备接受kafka发送的数据
func Init(address string, chansize, channums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogData, chansize)
	// 启动一个goroutine发送ES
	for i := 0; i < channums; i++ {
		go sendToES()
	}
	return
}

func SendToESChan(msg *LogData) {
	ch <- msg
}
func sendToES() {
	for {
		select {
		case msg := <-ch:
			// BodyJson传入一个可以被json格式化的类型,他会序列化一个
			put1, err := client.Index().Index(msg.Topic).Type("_doc").BodyJson(msg).Do(context.Background())
			if err != nil {
				// Handle error
				fmt.Println(err)
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}

}
