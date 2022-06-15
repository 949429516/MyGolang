package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	proto "socket_demo/tcp_nianbao/protoal"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		fmt.Println("接受client端发来的信息是:", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go process(conn)
	}
}
