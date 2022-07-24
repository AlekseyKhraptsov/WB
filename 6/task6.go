package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	signalTerm()
	// Ждем  Ctrl+C, который передается в горутину, для того, чтобы она завершалась

	withDeadLine()
	// закрывается по DeadLine
	withTimeout()
	// withTimeout получает сигнал о закрытии, спустя определенное время, после начала программы

	withCancel()
	// withCancel получает сигнал ошибки, и отправляет сигнал открывшимся горутинам о закрываться
}

func withDeadLine() {
	wg := sync.WaitGroup{}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	ch := make(chan string, 1)
	ch <- "Gophers on the attack"

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Close")
				close(ch)
				return
			default:
				mes := <-ch
				fmt.Println(mes)
				ch <- mes
			}
		}
	}()

	wg.Wait()
}

func withTimeout() {
	wg := sync.WaitGroup{}

	ctx := context.Background()
	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, d)
	defer cancel()

	ch := make(chan string, 1)
	ch <- "Gophers on the attack two"

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Close")
				close(ch)
				return
			default:
				mes := <-ch
				fmt.Println(mes)
				ch <- mes
			}
		}
	}()

	wg.Wait()
}

func longRunningProcess() error {
	time.Sleep(time.Second)
	return errors.New("Failed")
}

func withCancel() {
	wg := sync.WaitGroup{}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := longRunningProcess()
		if err != nil {
			cancel()
		}
	}()

	ch := make(chan string, 1)
	ch <- "Gophers on the attack three"

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Close")
				close(ch)
				return
			default:
				mes := <-ch
				fmt.Println(mes)
				ch <- mes
			}
		}
	}()

	wg.Wait()
}

func signalTerm() {
	wg := sync.WaitGroup{}

	stopSignal := make(chan os.Signal, 1)

	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)

	ch := make(chan string, 1)
	ch <- "Gophers on the attack"

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopSignal:
				fmt.Println("Close")
				close(ch)
				return
			default:
				mes := <-ch
				fmt.Println(mes)
				ch <- mes
			}
		}
	}()

	wg.Wait()
}
