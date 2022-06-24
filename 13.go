package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 0, 1
	fmt.Printf("a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
}
