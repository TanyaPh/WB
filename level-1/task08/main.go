package main

import "fmt"

func setBit(num int64, pos int, bitValue int) int64 {
    mask := int64(1) << pos 	// маска с установленным битом 
    if bitValue == 1 {
        return num | mask		// бит устанавливается в 1 
    } else {
        return num &^ mask		// бит устанавливается в 0
    }
}

func main() {
	var n int64 = 9223372036854775807
	fmt.Printf("%064b\n", n)

	for i := 0; i < 64; i += 2 {
		n = setBit(n, i, 0)
	}

	fmt.Printf("%064b\n", n)
}
