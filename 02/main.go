package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8, 10}
	done := make(chan struct{})
	for _, v := range nums {
		go func(v int) {
			fmt.Println(v * v)
			done <- struct{}{}
		}(v)
	}

	for range nums {
		<-done
	}
	close(done)
}
