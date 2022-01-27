package day03

import "fmt"

func AboutParameters() {
	// int
	fmt.Println("i=", i)
	modifyVars(i)
	fmt.Println("i=", i)

	// string
	fmt.Println("str=", str)
	modifyVars(str)
	fmt.Println("str=", str)

	// obj
	fmt.Println("o=", o)
	modifyVars(o)
	fmt.Println("o=", o)

	// arr
	fmt.Println("arr=", arr)
	modifyVars(arr)
	fmt.Println("arr=", arr)

	// slice
	fmt.Println("sli=", sli)
	fmt.Println("sli2=", sli2)
	modifyVars(sli)
	fmt.Println("sli=", sli)
	fmt.Println("sli2=", sli2)

	// map
	fmt.Println("hashmap=", hashMap)
	fmt.Println("hashmap2=", hashMap2)
	modifyVars(hashMap)
	fmt.Println("hashmap=", hashMap)
	fmt.Println("hashmap2=", hashMap2)

}

// define vars
var i int = 1

var str string = "string"

type obj struct {
	name string
}

var o = obj{
	name: "peter",
}

var arr [3]int = [3]int{3, 3, 3}

var sli []int = []int{1, 2, 3}
var sli2 = sli

var hashMap map[string]int = map[string]int{
	"age": 22,
}
var hashMap2 = hashMap

func modifyVars(any interface{}) {
	switch v := any.(type) {

	case int:
		v = 2

	case string:
		v = "new string"

	case obj:
		v = obj{
			name: "peter after modify",
		}

	case [3]int:
		v[0] = 222

	// reference
	case []int:
		v[0] = 2222

	case map[string]int:
		v["age"] = 222

	default:
		fmt.Println(v, "doesn't match")
	}
}
