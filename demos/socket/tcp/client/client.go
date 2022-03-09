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

// Client
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Client: net.Dial, failed with err: ", err)
		return
	}
	fmt.Println("Client: TCP connection connected now:", timex.GetNowString())
	defer conn.Close() // 关闭连接

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

		_, err = conn.Write([]byte(reqData)) // 向连接, 发送请求数据
		if err != nil {
			fmt.Println("Client: connection write data, failed with err: ", err)
			return
		}

		rspBuf := [512]byte{}
		n, err := conn.Read(rspBuf[:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Client: TCP connection, Server quit already!")
				break
			}
			fmt.Println("Client: connection receive data, failed with err: ", err)
			return
		}

		fmt.Println("Client: connection receive data from Server: ", string(rspBuf[:n]))
	}
	fmt.Println("Client: quit already:", timex.GetNowString())
}
