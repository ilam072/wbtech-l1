package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var workersNums int
	fmt.Scan(&workersNums)

	nums := make(chan int)

	// Используем context.WithCancel для завершения всех горутин
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	for i := 0; i < workersNums; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					// Когда получаем сигнал завершения - корректно выходим из горутины
					return
				case num, ok := <-nums:
					if !ok {
						return
					}
					fmt.Println(num)
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				// При отмене контекста закрываем канал, чтобы горутины могли корректно завершиться
				close(nums)
				return
			case nums <- rand.Int():
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	// Канал для получения сигнала завершения от ОС (например, CTRL + C)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	cancel()  // Отменяем контекст - отправляем горутина сигнал завершения
	wg.Wait() // Дожидаемся завершения всех горутин
	fmt.Println("program stopped")
}
