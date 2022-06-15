package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {
	defer conn.Close()
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}
func main() {
	//1.本地端口启动
	listener, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:8899 failed, err:", err)
		return
	}
	for {
		//2.等待建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		//3.与客户端通信
		go processConn(conn)
	}
}
