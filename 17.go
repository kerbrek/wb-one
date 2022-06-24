package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func binSearch(s []int, target int) int {
	left, right := 0, len(s)-1

	for left <= right {
		mid := (left + right) / 2
		if s[mid] < target {
			left = mid + 1
		} else if s[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	nums := []int{2, 7, 9, 12, 18, 24, 34, 45}
	fmt.Println("nums:", nums)
	num := 18
	fmt.Printf("%d in nums: %d\n", num, binSearch(nums, num))
	num = 42
	fmt.Printf("%d in nums: %d\n", num, binSearch(nums, num))
}
