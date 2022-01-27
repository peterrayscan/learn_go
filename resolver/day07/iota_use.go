package day07

import "fmt"

func IotaUse() {
	fmt.Println("IgEggs", IgEggs)
	fmt.Println("IgChocolate", IgChocolate)
	fmt.Println("IgNuts", IgNuts)
	fmt.Println("IgStrawberries", IgStrawberries)
	fmt.Println("IgShellfish", IgShellfish)

	fmt.Println("")
	fmt.Println(IgEggs | IgChocolate | IgShellfish) // got: 0001 0011, which is 19

	fmt.Println("")
	fmt.Println(x, y, z, k, p)

	fmt.Println("")
	fmt.Println(a, b, c)
}

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 0000 0001, 1
	IgChocolate                         // 1 << 1 which is 0000 0010, 2
	IgNuts                              // 1 << 2 which is 0000 0100, 4
	IgStrawberries                      // 1 << 3 which is 0000 1000, 8
	IgShellfish                         // 1 << 4 which is 0001 0000, 16
)

const (
	x = iota // 0
	_        // 1
	y        // 2
	z = "zz" // zz
	k        // zz
	p = iota // 5
)

const (
	a = "aa"
	b
	c
)
