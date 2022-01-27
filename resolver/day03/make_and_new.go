package day03

import (
	"fmt"
)

type sct struct {
	int32Value   int32
	stringValue  string
	float32Value float32
	boolValue    bool
}

func MakeAndNew() {
	MakeFunc()
	//NewFunc()
}

func MakeFunc() {
	// makeSlice()
	makeMap()
}

func makeSlice() {
	// intSlice := make([]int, 2)
	// printInfo(intSlice)
	// intSlice[0] = 123
	// intSlice[1] = 456
	// printInfo(intSlice)

	// sct
	sctSlice := make([]sct, 2)
	printInfo(sctSlice)
}

func makeMap() {
	intMap := make(map[int]int)
	printInfo(intMap)
	intMap[0] = 111
	intMap[2] = 222
	printInfo(intMap)

	// sct
	// sctMap := make(map[string]sct, 20)
	// printInfo(sctMap)
	// sctMap["test"] = sct{}
	// printInfo(sctMap)
}

func NewFunc() {
	newT()
}

func newT() {
	int32V := new(int32)
	fmt.Println(int32V)
	fmt.Println(*int32V)
	fmt.Println(&int32V)

	fmt.Println("")
	float32V := new(bool)
	fmt.Println(float32V)
	fmt.Println(*float32V)
	fmt.Println(&float32V)

	fmt.Println("")
	sctV := new(sct)
	fmt.Println(sctV)
	fmt.Println(*sctV)
	fmt.Println(&sctV)

}

func printInfo(i interface{}) {
	fmt.Printf("type: %T\n", i)
	fmt.Printf("addr: %p\n", i)

	switch any := i.(type) {
	case []int:
		fmt.Println("len: ", len(any), ", cap: ", cap(any))
		for index, value := range any {
			fmt.Printf("%T's [%v]: %v, addr: %p\n", any, index, value, &any[index])
		}

	case []sct:
		fmt.Println("len: ", len(any), ", cap: ", cap(any))
		for index, value := range any {
			fmt.Printf("%T's [%v]: %+v, addr: %p\n", any, index, value, &any[index])
		}

	case map[int]int:
		fmt.Println("len: ", len(any))
		for k, v := range any {
			temp := any[k]
			fmt.Printf("%T's key: %v, value: %v, value's addr: %p\n", any, k, v, &temp)
		}

	case map[string]sct:
		fmt.Println("len: ", len(any))
		for k, v := range any {
			temp := any[k]
			fmt.Printf("%T's key: %v, value: %+v, value's addr: %p\n", any, k, v, &temp)
		}

	default:
		fmt.Println("doesn't match")
	}
}
