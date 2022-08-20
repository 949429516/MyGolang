package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/*
客户端
1. 客户端一定要做好出入参数结构，需要关注的是结构内部的参数是否与rpc提供的对应入参需要对应出参数
2. 先创建client使用反法rpc.DiaHttp
3. client.Go(结构体名.方法名，入参，回参指针，chan 可以nil默认即可) 返回一个chan自行创建阻塞时间
4. client.Call(结构体名.方法名,入参,回参指针) 直接阻塞
*/
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
	client.Call("Server.Add", req, &res)
	fmt.Println(res)
}
