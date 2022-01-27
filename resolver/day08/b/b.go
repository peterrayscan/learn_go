package b

import "fmt"

func init() {
	fmt.Println("b package: b.go's init func")
}

func B() {
	fmt.Println("b package: b.go's B func")
}
