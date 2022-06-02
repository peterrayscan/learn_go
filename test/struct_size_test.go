package test

import (
	"fmt"
	"learn/tools/testx"
	"testing"
	"unsafe"
)

type NormalStruct struct {
	name string
}

type EmptyStruct struct{}

var EmptyStructWithVar = struct{}{}

func Test_(t *testing.T) {
	testx.RunFunc(func() {
		fmt.Println("size of NormalStruct:", unsafe.Sizeof(NormalStruct{}))
		fmt.Println("size of EmptyStruct:", unsafe.Sizeof(EmptyStruct{}))
		fmt.Println("size of EmptyStructWithVar:", unsafe.Sizeof(EmptyStructWithVar))
		fmt.Println("size of struct{}{}:", unsafe.Sizeof(struct{}{}))
	})
}
