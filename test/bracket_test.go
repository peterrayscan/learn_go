package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
)

func Test_BracketPriority(t *testing.T) {

	returnTrue := func() bool {
		fmt.Println("执行到这里, 说明if右边的括号, 提高了优先级, 先进行括号的运算")
		return true
	}

	testx.RunFunc(func() {
		fmt.Println("实际情况")
		if true || (returnTrue() == true) {
			return
		}
	})
}

func Test_BracketPriority2(t *testing.T) {
	returnTrueLeft := func() bool {
		fmt.Println("左边的执行了")
		return true
	}

	returnTrueRight := func() bool {
		fmt.Println("右边的执行了")
		return true
	}

	testx.RunFunc(func() {
		fmt.Println("实际情况")
		if returnTrueLeft() == true || (returnTrueRight() == true) {
			return
		}
	})

}
