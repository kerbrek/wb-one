package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func reverseWords(s string) string {
	splitted := strings.Split(s, " ")
	n := len(splitted)
	for i := 0; i < n/2; i++ {
		splitted[i], splitted[n-1-i] = splitted[n-1-i], splitted[i]
	}

	return strings.Join(splitted, " ")
}

func main() {
	str := "snow  dog sun"
	fmt.Printf("Original: %q\n", str)
	fmt.Printf("Reversed: %q\n", reverseWords(str))
}
