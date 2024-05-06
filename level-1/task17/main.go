package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2	//индекс среднего элемента
		if arr[mid] == target {			
			return mid					//если искомый элемент, возвращаем
		} else if arr[mid] < target {	
			left = mid + 1				//если меньше, двигаем левую границу
		} else {
			right = mid - 1				//иначе, двигаем левую границу
		}
	}

	return -1		//элемент не найден
}

func main() {
    arr := []int{26, 54, 34, 71, 35, 40, 70, 85, 54, 52}
    fmt.Println(binarySearch(arr, 40))
	fmt.Println(binarySearch(arr, 42))
}
