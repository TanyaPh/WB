package main

import (
	"fmt"
	"slices"
	"sort"
)

func quickSort(arr []int) []int {
    if len(arr) < 2 {       //длина меньше 2, сортировка не требуется
        return arr
    }

    mid := arr[len(arr)/2]  // средний элемент
    left := make([]int, 0)  // слайс для всех меньше mid
    right := make([]int, 0) // слайс для всех больше mid
	equal := make([]int, 0) // слайс для всех равных mid

    for _, v := range arr {
        if v < mid {
            left = append(left, v)      //добавляем элементы меньше mid
        } else if v > mid {
            right = append(right, v)    //добавляем элементы больше mid
        } else {
            equal = append(equal, v)    //добавляем элементы равные mid
        }
    }

    left = quickSort(left)   //повторяем с меньшими значениями, получаем результат
    right = quickSort(right) //повторяем с большими значениями, получаем результат

    return append(append(left, equal...), right...)     //объединяет отсортированные части, возвращаем
}

func main() {
    arr := []int{26, 54, 34, 71, 35, 40, 70, 85, 54, 52}
    fmt.Println(quickSort(arr))

    slices.Sort(arr)
    fmt.Println(arr)

    arr = []int{26, 54, 34, 71, 35, 40, 70, 85, 54, 52}
    sort.Ints(arr)
    fmt.Println(arr)
}
