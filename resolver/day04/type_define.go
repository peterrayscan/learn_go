package day04

import "fmt"

func TypeDefine() {
	var i int = 1
	fmt.Println("i=", i)

	var y YYY = i
	fmt.Println("y=", y)

	var x XXX = XXX(i)
	fmt.Println("x=", x)
	x.Method1()

}

type XXX int

type YYY = int

func (x XXX) Method1() {
	fmt.Println("x is xxx", x)
}

// not working:
// invalid receiver int (basic or unnamed type)
// func (y YYY) Method2() {
// 	fmt.Println("y is yyy", y)
// }
