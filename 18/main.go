package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	value int
}

func (c *counter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	c := counter{
		mu:    sync.Mutex{},
		value: 0,
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.value)
}
