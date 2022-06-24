package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.

type Counter struct {
	sync.RWMutex
	val int64
}

func (c *Counter) Inc() {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *Counter) Val() int64 {
	c.RLock()
	defer c.RUnlock()
	return c.val
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
