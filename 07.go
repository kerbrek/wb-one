package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	nums := [...]int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	squares := make(map[int]int)

	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			squares[num] = num * num
			mu.Unlock()
		}(n)
	}

	wg.Wait()
	fmt.Println(squares)
}
