package main

import (
	"fmt"
	"unicode"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна
// быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func areCharsUniq(s string) bool {
	chars := make(map[rune]bool)
	for _, r := range s {
		low := unicode.ToLower(r)
		if chars[low] {
			return false
		}
		chars[low] = true
	}

	return true
}

func main() {
	fmt.Println("Все ли символы в строке уникальные (регистронезависимо)?")
	s1 := "abcd"
	s2 := "aabcd"
	s3 := "abcdefAB"
	fmt.Println(s1, "--", areCharsUniq(s1))
	fmt.Println(s2, "--", areCharsUniq(s2))
	fmt.Println(s3, "--", areCharsUniq(s3))
}
