package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	nums := [...]int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	squares := [len(nums)]int{}

	for i, n := range nums {
		wg.Add(1)
		go func(idx int, num int) {
			defer wg.Done()
			mu.Lock()
			squares[idx] = num * num
			mu.Unlock()
		}(i, n)
	}

	wg.Wait()
	for _, s := range squares {
		fmt.Println(s)
	}
}
