package main

import "fmt"

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

type square struct {
	idx int
	val int
}

func main() {
	nums := [...]int{2, 4, 6, 8, 10}
	squaresCh := make(chan square)

	for i, n := range nums {
		go func(idx int, num int) {
			squaresCh <- square{idx, num * num}
		}(i, n)
	}

	squares := [len(nums)]int{}
	for range nums {
		s := <-squaresCh
		squares[s.idx] = s.val
	}

	for _, s := range squares {
		fmt.Println(s)
	}
}
