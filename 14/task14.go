package main

import (
	"fmt"
)

func main() {
	var test interface{}
	test = "Good day"
	checkType(test)
	test = 1
	checkType(test)
	test = true
	checkType(test)
}

func checkType(i interface{}) {

	switch i := i.(type) {
	default:
		fmt.Printf("unexpected type %T\n", i)
	case int:
		fmt.Printf("type %T\n", i)
	case string:
		fmt.Printf("type %T\n", i)
	case bool:
		fmt.Printf("type %T\n", i)
	case chan interface{}:
		fmt.Printf("type %T\n", i)
	}

}
