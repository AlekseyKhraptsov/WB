package main

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// отправляем числа в канал.
func writer(ctx context.Context, ch chan<- int) {
	for i := 0; true; i++ { //бесконечный цикл
		select {
		default:
			ch <- i
			time.Sleep(time.Second) // задержка для наглядности
		case <-ctx.Done():
			close(ch) // закрытие канала
			return
		}
	}
}

// читаем из канала до его закрытия.
func reader(ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		fmt.Printf("%d\n", num)
	}
	fmt.Println()
	wg.Done()
}

func main() {
	// создаём контекст с завершением через N секунд.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	wg := new(sync.WaitGroup)
	ch := make(chan int, 1)
	wg.Add(1)
	go writer(ctx, ch)
	go reader(ch, wg)
	wg.Wait()
}
