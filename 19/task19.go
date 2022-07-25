package main

import (
	"fmt"
	"log"
)

func main() {
	var messageWord string
	_, err := fmt.Scan(&messageWord)
	if err != nil {
		log.Fatal(err)
	}

	var newWord string

	for _, i := range messageWord {
		newWord = string(i) + newWord
	}

	fmt.Println(newWord)
}
