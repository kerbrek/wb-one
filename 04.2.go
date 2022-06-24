package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и выводят
// в stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ
// завершения работы всех воркеров.

func doWork(wg *sync.WaitGroup, data <-chan any, id int) {
	defer wg.Done()

	for item := range data {
		fmt.Printf("%d: %v\n", id, item)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Work is done (%d)\n", id)
}

func main() {
	var workersNum int
	flag.IntVar(&workersNum, "n", runtime.NumCPU(), "number of workers")
	flag.Parse()

	wg := sync.WaitGroup{}
	data := make(chan any, workersNum)

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go doWork(&wg, data, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	i := 0
	for {
		select {
		case data <- i:
			i++
		case <-sigChan: // ждем получения сигнала os.Interrupt
			// завершаем итерации воркеров с помощью закрытия канала с данными
			close(data)
			// ждем завершения воркеров
			wg.Wait()
			fmt.Println("Done")
			return
		}
	}
}
