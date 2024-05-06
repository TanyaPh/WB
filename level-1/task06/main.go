package main

import (
	"context"
	"fmt"
	"time"
)

//остановка с помощью канала
func stopMethod01() {
	fmt.Println("Start method 1")
	stop := make(chan bool)

    go func() {
        for {
            select {
            default:
				fmt.Println("Work")
            case <-stop:		//завершение работы
                return
            }
        }
    }()

	time.Sleep(2 * time.Second)
	stop <- true		//сигнал остановки
	fmt.Println("Stop")
	
	time.Sleep(2 * time.Second)
	fmt.Println("It was method 1")
}

//остановка с помощью контекста
func stopMethod02() {
	fmt.Println("Start method 2")
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		for {
			select {
			default:
				fmt.Println("Work")
			case <-ctx.Done():		//завершение работы
				return
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()		//сигнал остановки
	fmt.Println("Stop")
	
	time.Sleep(2 * time.Second)
	fmt.Println("It was method 2")
}

func main() {
	stopMethod01()
	stopMethod02()
}
