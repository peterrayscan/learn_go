package tcp

import (
	"fmt"
	"learn/demos/socket/tcp/server"
)

func RunTcp() {
	fmt.Println("server running...")
	server.Run()
}
