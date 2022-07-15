package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd watch
// use etcd/clientv3
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	//watch操作,用来改变通知
	if err != nil {
		fmt.Println("connect to etcd failed, err:", err)
		return
	}
	fmt.Println("connect etcd success")
	defer cli.Close()
	// watch

	// 派一个哨兵一直监视liqinyan的变换(新增、修改、删除)
	ch := cli.Watch(context.Background(), "liqinyan")

	// 只要变化这个通道就能获取到值
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v,key:%v,value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}
}
