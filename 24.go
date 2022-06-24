package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y
// и конструктором.

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (a Point) Distance(b Point) float64 {
	return math.Sqrt(math.Pow(b.X()-a.X(), 2) + math.Pow(b.Y()-a.Y(), 2))
}

func main() {
	a, b := NewPoint(2, 3), NewPoint(-3, -2)
	fmt.Println(a.Distance(b))
}
