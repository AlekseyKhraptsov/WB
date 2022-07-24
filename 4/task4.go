package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
*/

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run 4/task4.go <number of workers>")
		return
	}
	numberOfWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("'%s' is not a number", os.Args[1])
		return
	}
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	data := make(chan int, numberOfWorkers)

	wg.Add(numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		go func(id int) {
			defer wg.Done()

			for n := range data {
				fmt.Println("worker id:", id)
				fmt.Println("data:", n)
			}
			// второй вариант - использование ctx.Done() или просто done := make(chan struct{})
			// но мне этот вариант кажется более громоздким
		}(i)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-shutdown:
			close(data)
			wg.Wait()
			fmt.Println("all workers finished")
			return
		default:
			data <- rand.Int()
		}
	}
}
