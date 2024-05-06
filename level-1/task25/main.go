package main

import (
	"fmt"
	"time"
)

func sleep01(d time.Duration) {
    <-time.After(d)				//блокировка канала до истечения времени
}

func sleep02(d time.Duration) time.Time {
    ticker := time.Tick(d)		//возвращает канал времени, получим текущее время с заданным интервалом
    for done := range ticker {
        return done
    }
    return time.Now()
}


func main() {
	fmt.Println(time.Now().Second())
	sleep01(4 * time.Second)
	fmt.Println(time.Now().Second())
	sleep02(4 * time.Second)
	fmt.Println(time.Now().Second())
}
