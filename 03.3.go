package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов
// с использованием конкурентных вычислений.

func main() {
	nums := [...]int64{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	sum := int64(0)

	for _, n := range nums {
		wg.Add(1)
		go func(num int64) {
			defer wg.Done()
			atomic.AddInt64(&sum, num*num)
		}(n)
	}

	wg.Wait()
	fmt.Println(sum)
}
