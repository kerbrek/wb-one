package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func reverseWords(s string) string {
	res := []string{}
	const (
		outside = iota // находимся вне слова
		inside         // находимся внутри слова
	)
	state := outside
	b := strings.Builder{}

	for _, r := range s {
		switch state {
		case inside:
			if r == ' ' || r == '\t' {
				state = outside
				res = append(res, b.String())
				b.Reset()
			}
			b.WriteRune(r)
		case outside:
			if r != ' ' && r != '\t' {
				state = inside
				res = append(res, b.String())
				b.Reset()
			}
			b.WriteRune(r)
		}
	}
	res = append(res, b.String())

	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return strings.Join(res, "")
}

func main() {
	str := "snow\tdog sun\t"
	fmt.Printf("Original: %q\n", str)
	fmt.Printf("Reversed: %q\n", reverseWords(str))
}
