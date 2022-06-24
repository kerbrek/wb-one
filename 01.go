package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов). Реализовать
// встраивание методов в структуре Action от родительской структуры Human
// (аналог наследования).

type Human struct {
	Name string
	age  int
}

func (h Human) name() string {
	return h.Name
}

func (h Human) Age() int {
	return h.age
}

// struct embedding
// https://go.dev/doc/effective_go#embedding
// https://go.dev/ref/spec#Struct_types
type Action struct {
	Human
}

func main() {
	a := Action{Human{"Alex", 50}}
	fmt.Println("name method:", a.name())
	fmt.Println("Age method:", a.Age())
	fmt.Println("Name field:", a.Name)
	fmt.Println("age field:", a.age)
}
