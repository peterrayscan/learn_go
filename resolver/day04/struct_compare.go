package day04

import (
	"fmt"
)

func StructCompare() {
	p1 := person{
		name: "peter",
		age:  22,
	}

	p2 := person{
		name: "peter",
		age:  22,
	}

	// same person struct, but different assign order
	p3 := person{
		age:  22,
		name: "peter",
	}

	fmt.Println("(p1 == p2) is", (p1 == p2))
	fmt.Println("(p1 == p3) is", (p1 == p3))

	// compile fail
	// p4 := personWithDifferentFieldOrder{
	// 	name: "peter",
	// 	age:  22,
	// }
	// fmt.Println("(p1 == p4) is", (p1 == p4))
}

type person struct {
	name   string
	age    int
	grades [3]int
}

type personWithDifferentFieldOrder struct {
	name   string
	grades [3]int
	age    int
}
