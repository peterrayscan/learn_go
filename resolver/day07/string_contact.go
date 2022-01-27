package day07

import (
	"fmt"
	"strings"
)

func StringContact() {
	method1()
	method2()
	method3()
	method4()
	method5()
	method6()
	method7()
}

func method1() {
	str := "xxx" + "yyy"
	fmt.Println("str is", str)
}

func method2() {
	str := `xxx` + `yyy`
	fmt.Println("str is", str)
}

// with next line
func method3() {
	str := "xxx" +
		"yyy"
	fmt.Println("str is", str)
}

// with next line
func method4() {
	str := `
	xxx + yyy,hhhh
	`
	fmt.Println("str is", str)
}

func method5() {
	str := strings.Join([]string{"XXX", "YYY"}, "a")
	fmt.Println("str is", str)
}

func method6() {
	str := fmt.Sprintf("xxx is type: %T", 111)
	fmt.Println("str is", str)
}

func method7() {
	str := "xxx" + `yyy`
	fmt.Println(str)
}
