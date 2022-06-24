package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	i := 2

	fmt.Println("с сохранением порядка")

	nums := []int{2, 4, 6, 8, 10}
	fmt.Println("before:", nums)

	nums = append(nums[:i], nums[i+1:]...)

	fmt.Println("after :", nums)

	fmt.Println("без сохранения порядка")

	nums = []int{2, 4, 6, 8, 10}
	fmt.Println("before:", nums)

	nums[i] = nums[len(nums)-1]
	nums = nums[:len(nums)-1]

	fmt.Println("after :", nums)
}
