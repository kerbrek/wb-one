package main

import "fmt"

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
// из массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func gen(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func doubles(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			out <- num * 2
		}
	}()

	return out
}

func main() {
	nums := [...]int{2, 4, 6, 8, 10}

	for d := range doubles(gen(nums[:])) {
		fmt.Println(d)
	}
}
