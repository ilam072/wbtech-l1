package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Условие
	stop := false
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for !stop {
			fmt.Println("goroutine working...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("goroutine stopped by condition")
	}()

	time.Sleep(6 * time.Second)
	stop = true
	wg.Wait()
}

func main() {
	// Канал
	stop := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				fmt.Println("goroutine stopped by channel")
				return
			default:
				fmt.Println("goroutine working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(6 * time.Second)
	stop <- struct{}{}
	wg.Wait()
}

func main() {
	// context.WithCancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			defer wg.Done()
			select {
			case <-ctx.Done():
				fmt.Println("goroutine stopped by context with cancel")
				return
			default:
				fmt.Println("goroutine working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(6 * time.Second)
	cancel()
	wg.Wait()
}

func main() {
	// context.WithTimeout
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine stopped by timeout")
				return
			default:
				fmt.Println("goroutine working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	wg.Wait()
}

func main() {
	// runtime.Goexit()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("goroutine stopped by runtime.Goexit()")
			wg.Done()
		}()
		for {
			fmt.Println("goroutine working...")
			time.Sleep(5 * time.Second)
			runtime.Goexit()
		}
	}()
	wg.Wait()
}

func main() {
	// Закрытие канала
	stop := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				fmt.Println("goroutine stopped by channel closure")
				return
			default:
				fmt.Println("goroutine working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(6 * time.Second)
	close(stop)
	wg.Wait()
}
