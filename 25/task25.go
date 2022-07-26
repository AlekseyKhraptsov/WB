package main

// Задание 25
// Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func sleep(x time.Duration) {
	<-time.After(x)
}

func main() {
	x := time.Second * 3
	fmt.Printf("Sleeping for %v...", x)
	sleep(x)
	fmt.Println("time out!")

}
