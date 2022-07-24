package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	//Для конкурентной записи в map создаем RWMutex
	mux := sync.RWMutex{}

	newMap := make(map[int]string)

	ch := make(chan int)

	for i := 1; i < 6; i++ {
		go func(i int) {
			mux.Lock()
			newMap[i] = strconv.Itoa(i) + " step"
			mux.Unlock()
			ch <- i
		}(i)
	}

	for i := 0; i < 5; i++ {
		x := <-ch
		fmt.Println("Key: ", x, "Value: ", newMap[x])
	}
}
