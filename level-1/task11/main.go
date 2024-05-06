package main

import "fmt"

func intersection(m1, m2 map[int]struct{}) map[int]struct{} {
	m := map[int]struct{}{}

	for k := range m1 {
		if _, ok := m2[k]; ok {			
			m[k] = struct{}{}		// приналичии значение из m1 в m2 добавляем значение
		}
	}

	return m
}

func main() {
	m1 := map[int]struct{}{27:{}, 45:{}, 66:{}, 53:{}, 72:{}, 67:{}, 62:{}}
	m2 := map[int]struct{}{60:{}, 96:{}, 72:{}, 22:{}, 39:{}, 66:{}}

	fmt.Println(intersection(m1, m2))
}
