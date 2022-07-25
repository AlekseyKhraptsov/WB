package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	mass := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chanel1 := make(chan int, 10)
	chanel2 := make(chan int, 10)
	//проходим по массиву, пишем данные в chanel1
	for _, x := range mass {
		chanel1 <- x
	}
	close(chanel1)
	//читаем данные из chanel1 и передаем квадрат в chanel2
	wg.Add(1)
	go func() {
		for x := range chanel1 {
			chanel2 <- x * x
		}
		close(chanel2)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range chanel2 {
			fmt.Println(v)
		}
		wg.Done()
	}()

	wg.Wait()
}
