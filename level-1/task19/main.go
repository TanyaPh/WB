package main

import (
	"fmt"
	"strings"
)

//обмен рунами с начала и конца
func invertStr01(s string) string{
	str := []rune(s)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	return string(str)
}

// Используя Builder добавляем по руне с конца
func invertStr02(s string) string{
    var str strings.Builder
	runes := []rune(s)

	for i := len(runes)-1; i >= 0; i-- {
		str.WriteRune(runes[i])
	}

    return str.String()
}

// добавляет каждый следующий символ в начало строки 
func invertStr03(s string) string{
	var str string

	for _, v := range s {
		str = string(v) + str
	}

	return str
}

func main() {
	strs := [...]string{"The quick brown fox jumps over the lazy dog!",
						"Съешь ещё этих мягких французских булок да выпей чаю",
						"天地玄黄 宇宙洪荒 日月盈昃 辰宿列张"}
	for _, v := range strs {
		fmt.Println(v)
		fmt.Println(invertStr01(v))
		fmt.Println(invertStr02(v))
		fmt.Println(invertStr03(v))
	}
}
