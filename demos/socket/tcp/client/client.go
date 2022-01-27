package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端
func Run() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("net.Dial failed, err: ", err)
		return
	}
	defer conn.Close() // 关闭连接
	fmt.Println("connected!")
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("please input: ")
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			fmt.Println("quit already")
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err: ", err)
			return
		}
		fmt.Println("client端: 收到server端，返回的数据=", string(buf[:n]))
	}
}
