package etcd

import (
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化etcd的方法
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, err:", err)
		return
	}
	return
}

// 从etcd中根据key获取配置项目
func GetConf(key string) (logEntry []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntry)
		if err != nil {
			fmt.Println("UNmarshal etcd value failed, err:", err)
			return
		}
	}
	return
}
