package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
)

func Test_RunHandler(t *testing.T) {
	testx.RunFunc(func() {

		// 对于不同的参数，都动态执行
		handler1 := func() testx.RunHandler {
			return func() {
				fmt.Println("a is", 0)
			}
		}
		testx.Run(handler1())

		handler2 := func(a int) testx.RunHandler {
			return func() {
				fmt.Println("a is", a)
			}
		}
		testx.Run(handler2(2))

	})

}
