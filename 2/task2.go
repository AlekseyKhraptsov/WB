package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов
//чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	numbers := []int{2, 4, 6, 7, 8, 10}
	// Зелаем группу ожидания
	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)
		// передаём число в качестве аргумента горутине
		go func(number int) {
			defer wg.Done()

			fmt.Println(number * number)
		}(number)
	}
	//ждем окончания
	wg.Wait()
}
