package playground

import (
	"fmt"
	"strconv"
)

func ShortAssign1() {
	var err error
	fmt.Printf("%T: %p\n", err, &err)

	// 同一作用域内， data 声明变量，err 重新赋值
	data, err := strconv.Atoi("x")
	fmt.Printf("%T: %p\n", err, &err)

	if err == nil {
		fmt.Println("Atoi success:", data)
	}
}

func ShortAssign2() {
	var data int
	fmt.Printf("%v: %p\n", data, &data)

	// 同一作用域内， data 重新赋值，err 声明变量
	data, err := strconv.Atoi("x")
	fmt.Printf("%v: %p\n", data, &data)

	if err == nil {
		fmt.Println("data:", data)
	}
}

func ShortAssign3() {
	var err error
	fmt.Printf("%T: %p\n", err, &err)

	if true {
		// if的子作用域内， data和err 都被声明变量
		// err由于外面声明过，此处是被再次声明，所以是 变量名一样，变量存储的 地址 不一样
		data, err := strconv.Atoi("x")
		fmt.Printf("%T: %p\n", err, &err)

		if err == nil {
			fmt.Println("Atoi success:", data)
		}
	}

	// if的子作用域结束，err变量 在子作用域的地址被丢弃，重新恢复为 外层作用域的地址
	fmt.Printf("%T: %p\n", err, &err)
}
