package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
)

var (
	// 全局ES客户端
	client *elastic.Client
)

// 初始化ES,准备接受kafka发送的数据
func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return
	}
	fmt.Println("connect to es success")
	return
}

func SendToES(index string, data interface{}) error {

	put1, err := client.Index().Index(index).Type("_doc").BodyJson(data).Do(context.Background())
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return err
}
