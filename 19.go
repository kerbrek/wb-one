package main

import (
	"fmt"
	"os"
	"strings"
)

// Разработать программу, которая переворачивает подаваемую на вход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

// Не учитывает положение комбинируемых символов
// https://ru.wikipedia.org/wiki/Комбинируемый_символ
// https://stackoverflow.com/a/62743214/6475258
func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: Pass an argument")
		os.Exit(1)
	}

	fmt.Println(reverse(strings.Join(os.Args[1:], " ")))
}
