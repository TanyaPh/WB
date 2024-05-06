package main

import "fmt"

func toSet(s []string) map[string]int {
	m := map[string]int{}

	for _, v := range s {
		m[v]++			// кладем значение в мапу сохраняя колличество
	}

	return m
}

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(toSet(s))
}
