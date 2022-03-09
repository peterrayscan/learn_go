package playground

import "fmt"

func SwitchCase(str string) {
	switch str {
	case "x":
		fmt.Println("in x")
	case "y": // do nothing
	default:
		fmt.Println("in default")
	}
}
