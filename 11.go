package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func intersect(setA, setB map[int]struct{}) map[int]struct{} {
	intersection := make(map[int]struct{})

	min, max := setA, setB
	if len(setA) > len(setB) {
		min, max = setB, setA
	}

	for el := range min {
		if _, ok := max[el]; ok {
			intersection[el] = struct{}{}
		}
	}

	return intersection
}

func main() {
	setA := map[int]struct{}{1: {}, 4: {}, 7: {}}
	setB := map[int]struct{}{4: {}, 6: {}, 7: {}, 9: {}}
	fmt.Println(intersect(setA, setB))
}
