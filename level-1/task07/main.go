package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Method01 struct {
	m map[string]int
	sync.Mutex
}

func newMethod01() *Method01 {
	return &Method01 {
		m: map[string]int{},
	}
}

//использую мьютекс
func (m *Method01) insert(k string, v int) {
	m.Lock()			// блокировка доступа
	defer m.Unlock()	// отложенная разблокировка доступа
	m.m[k] = v			// доспут к ресурсу
}

func insertIntoMapMethod01() {
	m := newMethod01()

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			m.insert(strconv.Itoa(i), i)
		}(i)
	}

	wg.Wait()
	for k, v := range m.m {
		fmt.Println(k, v)
	}
}

// использую метод Store у sync.Map
func insertIntoMapMethod02() {
	m := sync.Map{}

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			m.Store(strconv.Itoa(i), i) // гарантирует эксклюзивный доступ
		}(i)
	}

	wg.Wait()
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}

func main() {
	insertIntoMapMethod01()
	insertIntoMapMethod02()
}
