package main

import (
	"fmt"
	"sync"
)

type storage struct {
	mu    sync.RWMutex
	users map[int]string
}

func newStorage() storage {
	return storage{
		mu:    sync.RWMutex{},
		users: make(map[int]string),
	}
}

func main() {
	users := []string{"Dima", "Ivan", "John", "Adam", "Michael", "Jacob"}
	m := newStorage()
	wg := &sync.WaitGroup{}

	for i, user := range users {
		wg.Add(1)
		go func(i int, user string) {
			defer wg.Done()
			m.mu.Lock()
			defer m.mu.Unlock()
			m.users[i] = user
		}(i, user)
	}
	wg.Wait()

	fmt.Println(m.users)
}
