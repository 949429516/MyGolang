package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

/*RPC
1. Go的RPC只支持go写的系统
2. GO的RPC有以下要求:
	2.1 函数首字母必须大写
	2.2 必须只有两个参数 第一个是接收的参数，第二个是返回给客户端的参数且必须为指针类型
	2.3 还要有一个返回值error
	func (t *T) MethodName(argType T1, replyType *T2) (err error){}
3. 服务端
rpc.Regist(new(符合rpc的结构体))
rpc.HandleHttp()借用http协议还作为rpc整体
net.Listen("tcp",":1234")创建一个listen
http.Serve(l,nil)启动
*/
type Server struct {
}
type Req struct {
	NumOne int
	NumTwo int
}
type Res struct {
	Num int
}

func (s *Server) Add(req Req, res *Res) error {
	time.Sleep(time.Second * 3)
	res.Num = req.NumOne + req.NumTwo
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("failed")
	}
	http.Serve(l, nil)
}
