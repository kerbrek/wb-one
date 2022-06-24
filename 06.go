package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	// С помощью отдельного канала
	wg0 := sync.WaitGroup{}
	stop := make(chan bool)

	wg0.Add(1)
	go func() {
		defer wg0.Done()
		for {
			select {
			case <-stop:
				return
			default:
				//
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	stop <- true
	wg0.Wait()
	fmt.Println("Done")

	// С помощью закрытия отдельного канала (бродкастинг)
	wg1 := sync.WaitGroup{}
	done := make(chan bool)

	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for {
			select {
			case <-done:
				return
			default:
				//
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	close(done)
	wg1.Wait()
	fmt.Println("Done")

	// С помощью закрытия канала, по которому горутина итерируется
	wg2 := sync.WaitGroup{}
	data := make(chan bool)

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for range data {
			//
		}
	}()

	time.Sleep(100 * time.Millisecond)
	close(data)
	wg2.Wait()
	fmt.Println("Done")

	// По таймауту
	wg3 := sync.WaitGroup{}

	wg3.Add(1)
	go func(timeout time.Duration) {
		defer wg3.Done()
		timeoutExceeded := time.After(timeout)
		for {
			select {
			case <-timeoutExceeded:
				return
			default:
				//
			}
		}
	}(100 * time.Millisecond)

	wg3.Wait()
	fmt.Println("Done")

	// С помощью пакета context
	wg4 := sync.WaitGroup{}
	c, cancel := context.WithCancel(context.Background())

	wg4.Add(1)
	go func(ctx context.Context) {
		defer wg4.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				//
			}
		}
	}(c)

	time.Sleep(100 * time.Millisecond)
	cancel()
	wg4.Wait()
	fmt.Println("Done")

	// По таймауту с помощью пакета context
	wg5 := sync.WaitGroup{}
	tc, tcancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer tcancel()

	wg5.Add(1)
	go func(ctx context.Context) {
		defer wg5.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				//
			}
		}
	}(tc)

	wg5.Wait()
	fmt.Println("Done")

	// С помощью общей переменной
	wg6 := sync.WaitGroup{}
	mu := sync.Mutex{}
	exit := false

	wg6.Add(1)
	go func() {
		defer wg6.Done()
		for {
			mu.Lock()
			end := exit
			mu.Unlock()
			if end {
				return
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	exit = true
	mu.Unlock()
	wg6.Wait()
	fmt.Println("Done")

	// С помощью общей переменной и пакета sync/atomic
	wg7 := sync.WaitGroup{}
	quit := int32(0)

	wg7.Add(1)
	go func() {
		defer wg7.Done()
		for {
			if atomic.LoadInt32(&quit) > 0 {
				return
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	atomic.StoreInt32(&quit, 1)
	wg7.Wait()
	fmt.Println("Done")
}
