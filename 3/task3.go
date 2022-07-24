package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов с использованием конкурентных вычислений.

func main() {
	numbers := []int64{2, 4, 6, 7, 8, 10}
	//Создаем группу ожидания
	var wg sync.WaitGroup

	var sum int64

	for _, number := range numbers {
		wg.Add(1)
		//передаем число
		go func(number int64) {
			defer wg.Done()
			//суммируем результат
			atomic.AddInt64(&sum, number*number)
		}(number)
	}
	wg.Wait()

	fmt.Println(sum)
}
