package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0

	wg := sync.WaitGroup{} //

	for i := 0; i < 50; i++ { // Запускаем 100 горутин
		wg.Add(1) // Добавляем 1 к счетчику ожидания завершения горутины
		go func() {

			defer wg.Done()                 // При завершении горутины отнимаем 1 из счетчика ожидания завершения
			count++                         // Инкриминируем счетчик
			fmt.Println("Рутина1: ", count) // Выводим счетчик для отслеживания конкуренции

		}()
		wg.Add(1)
		go func() {

			defer wg.Done()
			count++
			fmt.Println("Рутина2: ", count)

		}()
		wg.Wait() // Ждем завершения всех горутин
	}

	fmt.Println("Итого: ", count) // Выводим общее значение счетчика
}
