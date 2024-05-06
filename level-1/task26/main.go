package main

import (
	"fmt"
	"strings"
)

func checkUnique01(str string) bool {
	m := make(map[rune]bool)
	str = strings.ToLower(str)		// все к нижнему регистру

	for _, char := range str {
		if m[char] {
			return false			//при совпадении false
		}
		m[char] = true
	}

	return true
}

func checkUnique02(str string) bool {
	m := make(map[rune]struct{})
	str = strings.ToLower(str)		// все к нижнему регистру

	for _, char := range str {
		if _, ok := m[char]; ok {
			return false			//при совпадении false
		}
		m[char] = struct{}{}
	}

	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Println(checkUnique01(s1))
	fmt.Println(checkUnique02(s1))

	fmt.Println(checkUnique01(s2))
	fmt.Println(checkUnique02(s2))

	fmt.Println(checkUnique01(s3))
	fmt.Println(checkUnique02(s3))
}
