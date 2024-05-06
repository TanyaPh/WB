package main

import "fmt"

func swapMethod02(a, b int) (int, int) {
	return b, a
}

func swapMethod03(a, b int) (int, int) {
	b = a + b		// b увеличиваем на a
    a = b - a		// из (b+a) вычитает a, остается b
    b = b - a		// из (b+a) вычитает a(которое теперь равно b), остается а
	return a, b
}

func main() {
	a, b := 1, 2

	a, b = b, a
	fmt.Println(a, b)

	a, b = swapMethod02(a, b)
	fmt.Println(a, b)

	a, b = swapMethod03(a, b)
	fmt.Println(a, b)
}
