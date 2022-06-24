package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(70), nil)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(65), nil)

	sum := new(big.Int).Add(a, b)
	difference := new(big.Int).Sub(a, b)
	product := new(big.Int).Mul(a, b)
	quotient := new(big.Int).Div(a, b)

	fmt.Printf("a     = %v\n", a)
	fmt.Printf("b     = %v\n", b)
	fmt.Printf("a + b = %v\n", sum)
	fmt.Printf("a - b = %v\n", difference)
	fmt.Printf("a * b = %v\n", product)
	fmt.Printf("a / b = %v\n", quotient)
}
