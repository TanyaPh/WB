package main

import (
	"fmt"
	"sync"
)

// Расчитывает квадраты используя WaitGroup
func Method01(numbers []int) {
	var wg sync.WaitGroup
	wg.Add(len(numbers))
	
	for i := range numbers {
		go func(i int) {		// новая горутина
			defer wg.Done()
			square := numbers[i] * numbers[i] // в ней расчет
			fmt.Println(square)	// в ней вывод
		}(i)
	}

	wg.Wait()
}

// Расчитывает квадраты используя канал
func Method02(numbers []int) {
	ch := make(chan int, len(numbers))

	go calcResult(numbers, ch)	// расчет в новой горутине

	for i := 0; i < len(numbers); i++ {
		square := <-ch			//получение результата из канала
		fmt.Println(square)		// вывод результата
	}
}

func calcResult(numbers []int, ch chan int) {
	for _, num := range numbers {
		ch <- num * num			// отправка результата в канал
	}
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	Method01(numbers)
	Method02(numbers)
}
