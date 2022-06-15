package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//与server建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("dial tcp client on 127.0.0.1:8899 failed, err:", err)
		return
	}
	//发送数据
	var msg string
	// if len(os.Args) < 2 {
	// 	msg = "hello world!"
	// } else {
	// 	msg = os.Args[1]
	// }
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入:")
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
