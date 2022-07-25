package main

import "fmt"

func main() {
	a := 15
	b := 7

	a, b = b, a
	fmt.Println("Вариант 1: ", a, " ", b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("Вариант 1: ", a, " ", b)
}
