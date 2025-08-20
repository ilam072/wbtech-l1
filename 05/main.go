package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	nums := make(chan int)
	duration := time.Second * 6
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case n, ok := <-nums:
				if !ok {
					return
				}
				fmt.Println(n)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(nums)
				return
			case nums <- rand.Intn(1000):
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
