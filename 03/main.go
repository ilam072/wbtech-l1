package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var workersNums int
	fmt.Print("Введите количество воркеров: ")
	_, err := fmt.Scan(&workersNums)
	if err != nil || workersNums <= 0 {
		fmt.Println("Некорректный ввод, требуется положительное число")
		return
	}

	nums := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < workersNums; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for n := range nums {
				fmt.Printf("Worker %d: %d\n", id, n)
			}
		}(i + 1)
	}

	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				close(nums)
				return
			default:
				nums <- rand.Int()
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	fmt.Println("\nПрограмма завершает работу...")

	close(done) // сигнал для остановки генератора
	wg.Wait()   // ждём завершения всех воркеров
	fmt.Println("Все воркеры остановлены, программа завершена")
}
