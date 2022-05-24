package test

import (
	"fmt"
	"learn/tools/testx"
	"strconv"
	"testing"
)

func Test_VariableScope1(t *testing.T) {
	testx.RunFunc(func() {

		withSubScope := func() testx.RunHandler {
			return func() {
				var err error
				fmt.Printf("value: %v, addr: %p\n", err, &err)
				if true {
					_, err := strconv.Atoi("x")
					// 子作用域内，返回的err，会将外层声明过的变量err，指向一个新的地址
					fmt.Printf("value: %v, addr: %p\n", err, &err)
				}
				// 子作用域结束后，err再次指回外层作用域的地址，子作用域的地址失效回收
				// 所以在处理这种子作用域的err时，稍不注意就可能踩坑
				// 解决办法：
				// 1. 最好直接就在子作用域内执行return
				// 2. 不要使用相同变量名
				fmt.Printf("value: %v, addr: %p\n", err, &err)
			}
		}
		testx.Run(withSubScope())

	})
}
