package main

import "fmt"

func main() {
	things := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(setOf(things))
}

func setOf(seq []string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, v := range seq {
		set[v] = struct{}{}
	}

	return set
}
