package main

import "fmt"

func remove(old []int, i int) []int {
	new := make([]int, len(old)-1)

	copy(new, old[:i])			//копируем слайс до удаляемого элемента
	copy(new[i:], old[i+1:])	//копируем слайс после удаляемого элемента

	return new					//возвращаем новый слайс
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	sl = remove(sl, 2)
	fmt.Println(sl)
}
