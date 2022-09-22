package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	micro "github.com/asim/go-micro/v3"
	pb "wklgrpc.com/micor/02/proto"
)

type Example struct{}
type Foo struct{}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Print("收到Example.Call请求")
	if len(req.Name) == 0 {
		return errors.New("go micro.api.example no name")
	}
	rsp.Message = "Example.Call接收到了你的请求" + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *pb.EmptyRequest, rsp *pb.EmptyResponse) error {
	log.Print("收到Foo.Bar请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
		micro.Address("127.0.0.1:8080"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()
	err := pb.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		fmt.Println(err)
	}
	err = pb.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		fmt.Println(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
