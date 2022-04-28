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

// UDP client端
func main() {
	socketConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败, err:", err)
		return
	}

	defer socketConn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Client: please input: ")
		inputContent, err := inputReader.ReadString('\n') // 读取用户输入
		if err != nil {
			fmt.Println("Client: read from console, failed with err: ", err)
			return
		}

		reqData := strings.Trim(inputContent, "\r\n")
		if strings.ToUpper(reqData) == "Q" { // 如果输入q/Q, 就退出
			fmt.Println("Client: ready to quit")
			break
		}

		_, err = socketConn.Write([]byte(reqData)) // 向连接, 发送请求数据
		if err != nil {
			fmt.Println("Client: connection write data, failed with err: ", err)
			return
		}

		rspBuf := [512]byte{}
		n, remoteAddr, err := socketConn.ReadFromUDP(rspBuf[:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Client: TCP connection, Server quit already!")
				break
			}
			fmt.Println("Client: connection receive data, failed with err: ", err)
			return
		}

		fmt.Println("Client: connection receive data from Server:")
		fmt.Println("data:", string(rspBuf[:n]))
		fmt.Println("remote addr:", remoteAddr)
		fmt.Println("length:", n)
	}
	fmt.Println("Client: quit already:", timex.GetNowString())
}
