package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("time:", time.Now())
	fmt.Println("sleep 1 sec")
	sleep(1 * time.Second)
	fmt.Println("time:", time.Now())
}
