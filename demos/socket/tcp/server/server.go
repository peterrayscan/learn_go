package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"learn/tools/timex"
	"net"
	"os"
	"strings"
)

var connectCount int = 1

// Server
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Server: net.Listen, failed with err: ", err)
		return
	}

	go quit()

	for {
		tcpConn, err := listener.Accept() // 建立连接
		if err != nil {
			fmt.Println("Server: listener.Accept, failed with err: ", err)
			fmt.Println("Server: continue wait for next connection")
			continue
		}
		fmt.Println("Server: TCP connection connected now:", timex.GetNowString())
		fmt.Println("Server: TCP connection, try to connect for", connectCount, "times")
		connectCount++
		go process(tcpConn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func process(tcpConn net.Conn) {
	defer tcpConn.Close() // 关闭连接
	for {
		connReader := bufio.NewReader(tcpConn)
		var reqBuf [128]byte
		n, err := connReader.Read(reqBuf[:]) // 读取数据
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Server: TCP connection, Client quit already:", timex.GetNowString())
				break
			}
			fmt.Println("Server: read from client, failed with err: ", err)
			return
		}

		reqString := string(reqBuf[:n])
		fmt.Println("Server: connection receive data from Client: ", reqString)

		splitedString := strings.Split(reqString, "")        // 分割字符串
		joinedBlankSpace := strings.Join(splitedString, " ") // 合并字符串

		tcpConn.Write([]byte("合并后的输入是: " + joinedBlankSpace)) // 发送数据
	}
	fmt.Println("Server: TCP connection disconnected now:", timex.GetNowString())
}

// quit os
func quit() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Server: running now, press the Q/q to quit")
	for {
		inputContent, err := inputReader.ReadString('\n') // 读取用户输入
		if err != nil {
			fmt.Println("Server: read from console, failed with err: ", err)
			os.Exit(1)
		}
		cmd := strings.Trim(inputContent, "\r\n")
		if strings.ToUpper(cmd) == "Q" { // 如果输入Q/q, 就退出
			fmt.Println("Server: quit already:", timex.GetNowString())
			os.Exit(0)
		}
	}
}
