package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8, 10}

	results := make(chan int, len(nums))

	for _, v := range nums {
		go func(v int) {
			results <- v * v
		}(v)
	}

	for range nums {
		fmt.Println(<-results)
	}

	close(results)
}
