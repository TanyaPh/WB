package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func worker(ch <-chan int) {
	for {
		fmt.Println(<-ch) 	// читает и выводит данные из канала 
	}
}

func run(ctx context.Context, ch chan<- int) {
	for i := -400000; ;i++ {
		select {
			case <-ctx.Done():	// получен сигнал, завершаем работу
				close(ch)
				return
			default:
				ch <- i			// отправляем данные в канал
		}
	}
}

func main() {
	var n int			 // выбора количества воркеров 
	fmt.Scan(&n)

	// ожидаем Ctrl+C
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ch := make(chan int, n)
	for i := 0; i < n; i++ {
		go worker(ch)
	}

	run(ctx, ch)
}
