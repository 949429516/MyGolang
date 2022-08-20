package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	NumOne int
	NumTwo int
}
type Res struct {
	Num int
}

func main() {
	req := Req{1, 2}
	var res Res
	// 创建客户端连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	call := client.Go("Server.Add", req, &res, nil)
	for {
		select {
		case <-call.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("等你哦")
		}
	}
}
