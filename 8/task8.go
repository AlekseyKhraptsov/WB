package main

import (
	"fmt"
)

//Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBit(x int64, i int, bool2 bool) int64 {
	// x заданное число
	// i бит в котором необходимо установить 1 (в заданном числе)

	if bool2 == true {
		return x | (1 << i)
	}
	return x &^ (1 << i)
}

func main() {
	fmt.Println(setBit(100, 3, true))

}
