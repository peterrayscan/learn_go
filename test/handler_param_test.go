package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
)

func Test_RunHandler(t *testing.T) {

	testx.RunFunc(func() {

		// 对于函数不同的参数，都动态执行
		// 没有参数的情况
		handler1 := func() testx.RunHandler {
			return func() {
				fmt.Println("no param")
			}
		}
		testx.Run(handler1())

		// 有参数的情况
		handler2 := func(a int) testx.RunHandler {
			return func() {
				fmt.Println("got param: ", a)
			}
		}
		testx.Run(handler2(2))

	})

}
