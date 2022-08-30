package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	// 设置续期为5秒，超过5秒没有续期则失效
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	// 将k-v设置到etcd
	_, err = cli.Put(context.TODO(), "root", "admin", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
	// 若想一直有效则要设置自动续期
	ch, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c := <-ch
		fmt.Println(c)
	}
}

/*
执行后一直在续期
cluster_id:14841639068965178418 member_id:10276657743932975437 revision:18 raft_term:13
cluster_id:14841639068965178418 member_id:10276657743932975437 revision:18 raft_term:13
cluster_id:14841639068965178418 member_id:10276657743932975437 revision:18 raft_term:13
cluster_id:14841639068965178418 member_id:10276657743932975437 revision:18 raft_term:13
cluster_id:14841639068965178418 member_id:10276657743932975437 revision:18 raft_term:13
cluster_id:14841639068965178418 member_id:10276657743932975437 revision:18 raft_term:13

set etcdctl_api=3
./etcdctl get root
admin
root

程序停止5s后 ./etcdctl get root 命令获取不到
*/
