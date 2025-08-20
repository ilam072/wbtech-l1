package main

import (
	"fmt"
)

func RemoveAt[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		return slice
	}

	copy(slice[i:], slice[i+1:])
	var zero T
	slice[len(slice)-1] = zero
	return slice[:len(slice)-1]
}

func main() {
	s1 := []string{"A", "B", "C", "D", "E"}
	s1 = RemoveAt(s1, 2)
	fmt.Println(s1)

	s2 := []int{10, 20, 30, 40, 50}
	s2 = RemoveAt(s2, 0)
	fmt.Println(s2)
}
