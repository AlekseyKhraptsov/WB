package main

import (
	"fmt"
	"log"
)

func main() {
	randomSlice := []int{1, 2, 3, 4, 5}

	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Fatal(err)
	}

	scissors(randomSlice, i)
}

func scissors(randomSlice []int, i int) {
	var newSlice []int

	/* Записываем в новый срез все до введенного числа и все после*/
	newSlice = append(randomSlice[:i], randomSlice[i+1:]...)

	fmt.Println(newSlice)
}
