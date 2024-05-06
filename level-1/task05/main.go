package main

import (
	"context"
	"fmt"
	"time"
)

func publisher(ctx context.Context, ch chan<- int) {
	for i := -400000; ; i++ {
		select {
		case <-ctx.Done():		//время вышло
			close(ch)
			return
		default:
			ch <- i				//отправка данных
		}
	}
}

func consumer(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():		//время вышло
			return
		default:
			val := <-ch			//получение данных
			fmt.Println(val)
		}
	}
}

func main() {
	var n time.Duration
	fmt.Scan(&n)	

	// задает таймер на n
	ctx, cancel := context.WithTimeout(context.Background(), n * time.Second)
	defer cancel()

	ch := make(chan int)
	go publisher(ctx, ch)
	go consumer(ctx, ch)

	<-ctx.Done() 	//ждет завешения таймера
}
