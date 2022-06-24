package main

import "fmt"

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов
// с использованием конкурентных вычислений.

func main() {
	nums := [...]int{2, 4, 6, 8, 10}
	squares := make(chan int)

	for _, n := range nums {
		go func(num int) {
			squares <- (num * num)
		}(n)
	}

	sum := 0
	for range nums {
		sum += <-squares
	}

	fmt.Println(sum)
}
