package console

import (
	"fmt"
)

func Before() {
	fmt.Println("———————————————————————————————————————————————————————————————————————————————————————————")
	fmt.Println("")
	// fmt.Printf("%v\n\n", timex.GetNowString())
}

func After() {
	// fmt.Printf("\n%v\n", timex.GetNowString())
	fmt.Println("")
	fmt.Println("———————————————————————————————————————————————————————————————————————————————————————————")
}
