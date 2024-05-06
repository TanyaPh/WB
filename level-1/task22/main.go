package main

import (
	"fmt"
	"math/big"
)

func main() {
	//создаем числа
	a, ok := big.NewInt(0).SetString("400400400400400400400400400400400400", 10)
	if !ok {
		fmt.Println("big.Int is not created")
	}

	b, ok := big.NewInt(0).SetString("400400400400400400400400400400400", 10)
	if !ok {
		fmt.Println("big.Int is not created")
	}

	//опирации с числами
	fmt.Println(new(big.Int).Add(a, b))
	fmt.Println(new(big.Int).Div(a, b))
	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Println(new(big.Int).Sub(a, b))


	//создаем числа
	a = big.NewInt(-200200200200200200)
	b = big.NewInt(-200200200200200)

	//опирации с числами
	fmt.Println(new(big.Int).Add(a, b))
	fmt.Println(new(big.Int).Div(a, b))
	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Println(new(big.Int).Sub(a, b))
}
