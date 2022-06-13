package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
)

func Test_MakeFuncLenCap(t *testing.T) {
	testx.RunFunc(func() {
		// i := make([]int, 5)
		i := make([]int, 0, 5)
		printIntSlice(i)
		i = append(i, 1)
		printIntSlice(i)
	})
}

func printIntSlice(intSlice []int) {
	fmt.Println("value:", intSlice, ", len:", len(intSlice), ", cap:", cap(intSlice))
}
