package main

import (
	"fmt"
	"sync"
)

func Method01(nums []int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result int

	wg.Add(len(nums))
	for _, num := range nums {
		go func(n int) {
			defer wg.Done()
			mu.Lock()			
			result += n * n	// только одна горутина может выполнить в одинмомент времени
			mu.Unlock()
		}(num)
	}
	wg.Wait()

	return result
}

func Method02(nums []int) int {
	ch := make(chan int, len(nums))

	for i := range nums {
		go func(i int) {	// новая горутина
			n := nums[i]
			ch <- n * n		// отправка результата в канал
		}(i)
	}

	result := 0
	for i := 0; i < cap(ch); i++ {
		result += <-ch		// получение результатов из канала
	}

	return result
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(Method01(nums))
	fmt.Println(Method02(nums))
}
