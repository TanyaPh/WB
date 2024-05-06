package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

	
type Counter interface {
    Increment()
}

type Counter01 struct {
	value uint64
}

func (c *Counter01) Increment() {
	atomic.AddUint64(&c.value, 1)	//увеличение счетчика через атомик
}

type Counter02 struct {
	value uint
	mu    sync.Mutex
}

func (c *Counter02) Increment() {
	c.mu.Lock()			//блокировка доступа
	defer c.mu.Unlock()	//отложенная разблокировка доступа
	c.value++			//увеличение счетчика 
}


func worker(counter Counter) {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
}

func main() {
	counter01 := &Counter01{}
    worker(counter01)
	fmt.Println("Counter01:", counter01.value)

	counter02 := &Counter02{}
    worker(counter02)
	fmt.Println("Counter02:", counter02.value)
}
