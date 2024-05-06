package main

import "fmt"

//Не эфективное использование ресурсов, строка не будет освобождена из памяти до тех пор, пока подстрока существует.

func main() {
	v := createHugeString(1 << 10)
	byteStr:= make([]byte, 100)

	//копируем необходимые данные
	copy(byteStr, v[:100])
	justString := string(byteStr)

	//перестает узазывать на огромную строку
	v = ""

	fmt.Println(justString)
}
