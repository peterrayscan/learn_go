package server

import (
	"bufio"
	"fmt"
	"learn/demos/socket/tcp/client"
	"net"
	"strings"
)

// TCP server端

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("server端: 收到client端，发来的数据=", recvStr)
		splitedString := strings.Split(recvStr, "")
		joinedBlankSpace := strings.Join(splitedString, " ")
		conn.Write([]byte(joinedBlankSpace)) // 发送数据
	}
}

func Run() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("net.Listen failed, err: ", err)
		return
	}

	// 在 server端 成功启动后，再启动 client端
	go client.Run()

	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("listen.Accept failed, err: ", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
