package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

// type switch https://go.dev/tour/methods/16
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int, chan string, chan bool:
		fmt.Println("channel")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// https://stackoverflow.com/questions/38748098/golang-type-switch-how-to-match-a-generic-slice-array-map-chan
func doReflect(i interface{}) {
	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("channel")
	default:
		fmt.Printf("I don't know about type %T!\n", v.Interface())
	}
}

func main() {
	do(1)
	do("str")
	do(true)
	do(make(chan int))
	do(make(chan string))
	do(make(chan bool))
	do(make(chan float64))
	do([]int{})

	fmt.Print("\n")

	doReflect(1)
	doReflect("str")
	doReflect(true)
	doReflect(make(chan int))
	doReflect(make(chan string))
	doReflect(make(chan bool))
	doReflect(make(chan float64))
	doReflect([]int{})
}
