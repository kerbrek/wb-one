package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов
// с использованием конкурентных вычислений.

func main() {
	nums := [...]int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	sum := 0

	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			sum += num * num
			mu.Unlock()
		}(n)
	}

	wg.Wait()
	fmt.Println(sum)
}
