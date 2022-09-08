package main

import (
	"context"
	"fmt"
	"log"

	micro "github.com/asim/go-micro/v3"
	pb "wklgrpc.com/micor/01/proto"
)

// 声明结构体
type Hello struct{}

func (g *Hello) Info(ctx context.Context, req *pb.InfoRequest, rep *pb.InfoResponse) (err error) {
	rep.Msg = "你好" + req.Username
	return
}

func main() {
	// 1.得到服务端实例
	service := micro.NewService(
		// 设置微服务的服务名，用来访问
		// micro call hello Hello.Info {"username":"zhangsan"}
		micro.Name("hello"),
	)
	// 2.初始化实例
	service.Init()
	// 3.服务注册
	err := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
