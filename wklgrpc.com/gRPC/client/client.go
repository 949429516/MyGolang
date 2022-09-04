package main

// 1.连接服务端
// 2.实例化gRPC客户端
// 3.调用
import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	pb "wklgrpc.com/gRPC/proto"
)

func main() {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听异常", err)
	}
	defer conn.Close()
	// 2.实例化客户端
	client := pb.NewUserInfoServiceClient(conn)
	// 3.组装访问参数
	req := new(pb.UserRequest)
	req.Name = "zs"
	// 4.调用接口
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常", err)
	}
	fmt.Println("响应结果:", response)
}
