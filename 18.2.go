package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.

type Counter struct {
	val int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.val, 1)
}

func (c *Counter) Val() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	counter := &Counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Val())
}
