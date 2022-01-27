package day07

import "fmt"

// nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量
func NilAssign() {
	// pointer
	var intPtr *int
	intPtr = nil
	fmt.Println(intPtr)

	// channel
	var ch chan int
	ch = nil
	fmt.Println(ch)

	// func
	var funcVar func(s string) int
	funcVar = nil
	fmt.Println(funcVar)

	// map and slice
	var m map[int]int
	m = nil
	fmt.Println(m)

	var s []int
	s = nil
	fmt.Println(s)
}
