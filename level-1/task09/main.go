package main

import "fmt"

func stage01(out chan<- int, nums []int) {
	for _, v := range nums {
		out <- v		// отправка даннаых из масива в канал
	}
	close(out)
}

func stage02(in <-chan int, out chan<- int) {
	for n := range in {		// получение из канала
		out <- n * 2		// отправка результата в канал
	}
	close(out)
}

func stage03(in <-chan int) {
	for res := range in {	// получение из канала
		fmt.Println(res)	// вывод полученных данных
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(nums))
	ch2 := make(chan int, len(nums))

	go stage01(ch, nums)
	go stage02(ch, ch2)
	stage03(ch2)
}
