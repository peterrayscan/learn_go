package playground

import "fmt"

type temp struct {
	str string
}

func StructNilValue() {
	t1 := &temp{
		str: "xxx",
	}
	fmt.Println(t1.str)

	t2 := &temp{}
	fmt.Println(t2.str)
}
