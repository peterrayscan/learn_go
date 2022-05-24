package testx

import "fmt"

type RunHandler func()

func Run(handler RunHandler) {
	handler()
	fmt.Println("")
}
