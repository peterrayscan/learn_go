package main

import (
	"fmt"
	"learn/demos/socket/tcp"
	"learn/tools/console"
)

func main() {
	console.Clear()
	fmt.Println("it works, git remote remove origin, then git remote add oringin xxx")
	tcp.RunTcp()
}
