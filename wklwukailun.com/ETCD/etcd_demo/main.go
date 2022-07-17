package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd client put/get demo
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
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//_, err = cli.Put(ctx, "wangsenbao", "dsb")
	//value := `[{"path":"/opt/nginx.log","topic":"web_log"},{"path":"/usr/redis.log","topic":"redis_log"},{"path":"/usr/mysql.log","topic":"mysql_log"}]`
	value := `[{"path":"/usr/mysql.log","topic":"mysql_log"}]`

	_, err = cli.Put(ctx, "/logagent/collecct_config", value)
	cancel()
	if err != nil {
		fmt.Println("put to etcd failed, err:", err)
		return
	}

	// get
	// ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	// resp, err := cli.Get(ctx, "/xxx")
	// cancel()
	// if err != nil {
	// 	fmt.Println("get to etcd failed, err:", err)
	// 	return
	// }
	// for _, ev := range resp.Kvs {
	// 	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	// }
}
