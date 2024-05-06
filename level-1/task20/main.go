package main

import (
	"fmt"
	"strings"
)

func invertWords01(s string) string{
	words := strings.Fields(s)				//разбиваем строку на слова
	str := make([]string, len(words))

	for i, j := 0, len(words)-1; i <= j; i, j = i+1, j-1 {
		str[i], str[j] = words[j], words[i]		//меняем слова местами с концов строки к серидине
	}

	return strings.Join(str, " ")
}

//добовляем слова с конца через Builder
func invertWords02(s string) string{
    var str strings.Builder
	words := strings.Fields(s)

	for i := len(words)-1; i >= 0; i-- {
		str.WriteString(words[i])
		str.WriteString(" ")
	}

    return str.String()
}

//каждое последущее слово добавляем в начало строки
func invertWords03(s string) string{
	var str string
	words := strings.Fields(s)

	for _, v := range words {
		str = v + " " + str
	}

	return str
}

func main() {
	strs := [...]string{"The quick brown fox jumps over the lazy dog!",
						"Съешь ещё этих мягких французских булок да выпей чаю",
						"天地玄黄 宇宙洪荒 日月盈昃 辰宿列张"}

	for _, v := range strs {
		fmt.Println(v)
		fmt.Println(invertWords01(v))
		fmt.Println(invertWords02(v))
		fmt.Println(invertWords03(v))
	}
}
