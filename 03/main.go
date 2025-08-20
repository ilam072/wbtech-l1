package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var workersNums int
	fmt.Scan(&workersNums)

	nums := make(chan int)

	for i := 0; i < workersNums; i++ {
		go func() {
			for {
				fmt.Println(<-nums)
			}
		}()
	}

	go func() {
		for {
			nums <- rand.Int()
			time.Sleep(300 * time.Millisecond)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	fmt.Println("program stopped")
}
