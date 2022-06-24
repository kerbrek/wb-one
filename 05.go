package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения
// в канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

func main() {
	var seconds int
	flag.IntVar(&seconds, "t", 1, "timeout in seconds")
	flag.Parse()

	wg := sync.WaitGroup{}
	data := make(chan any)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for item := range data {
			fmt.Println(item)
			time.Sleep(100 * time.Millisecond)
		}

		fmt.Println("Work is done")
	}()

	timeout := time.Duration(seconds) * time.Second
	timeoutExceeded := time.After(timeout)

	i := 0
	for {
		select {
		case data <- i:
			i++
		case <-timeoutExceeded:
			close(data)
			wg.Wait()
			fmt.Println("Done")
			return
		}
	}
}
