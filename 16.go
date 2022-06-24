package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func partition(s []int) int {
	start, end := 0, len(s)-1
	pivot := s[end]
	i := start

	for j := start; j < end; j++ {
		if s[j] <= pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[end] = s[end], s[i]

	return i
}

func quickSort(s []int) {
	if len(s) < 2 {
		return
	}

	i := partition(s)
	quickSort(s[:i])
	quickSort(s[i+1:])
}

func main() {
	nums := [...]int{2, 34, 12, 7, 45, 9, 24, 18}
	fmt.Println("before:", nums)
	quickSort(nums[:])
	fmt.Println("after :", nums)
}
