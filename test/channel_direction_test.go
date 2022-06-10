package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
)

func send(cin chan<- string, s string) {
	cin <- s
}

func get(cin chan<- string, cout <-chan string) {
	cin <- <-cout
}

func Test_ChannelDirection(t *testing.T) {
	testx.RunFunc(func() {
		setter := make(chan string, 1)
		getter := make(chan string, 1)
		send(setter, "xxx")
		get(getter, setter)
		fmt.Println(<-getter)
	})
}
