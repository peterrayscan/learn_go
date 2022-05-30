package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
)

func Test_SliceAppend(t *testing.T) {
	testx.RunFunc(func() {
		intSlice1 := make([]int, 3, 5)
		printSlice(intSlice1)

		// assign to another one: from index 0
		intSlice2 := intSlice1[0:2]
		printSlice(intSlice2)

		// assign to another one: from index 1
		intSlice3 := intSlice1[1:3]
		printSlice(intSlice3)
	})
}

func printSlice(intSlice []int) {
	fmt.Println("len:", len(intSlice), ", cap:", cap(intSlice))
}
