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

// 需要收集日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`  // 日志存放的路径
	Topic string `json:"topic"` // 日志要发往kafka中的topic
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
			fmt.Println("Unmarshal etcd value failed, err:", err)
			return
		}
	}
	return
}

// etcd watch
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	// 只要变化这个通道就能获取到值
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v,key:%v,value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			// 通知taillog.tskMgr有新配置
			// 1.判断类型
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				// 如果是删除操作,手动传递空的
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					continue
				}
			}
			newConfCh <- newConf
		}
	}
}
