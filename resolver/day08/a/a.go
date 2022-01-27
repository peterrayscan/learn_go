package a

import (
	"fmt"
	"learn/resolver/day08/b"
)

func init() {
	fmt.Println("a package: a.go's init func")
}

func A() {
	fmt.Println("a package: a.go's A func")
	b.B()
}
