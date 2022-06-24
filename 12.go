package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree). Создать для нее
// собственное множество.

func makeSet(sequence []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, str := range sequence {
		set[str] = struct{}{}
	}
	return set
}

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(makeSet(seq))
}
