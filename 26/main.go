package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isUnique("aAbcdt"))
	fmt.Println(isUnique("abCdtt"))
	fmt.Println(isUnique("a"))
	fmt.Println(isUnique("qwertyi"))
}

func isUnique(s string) bool {
	counter := make(map[rune]int)

	for _, l := range strings.ToLower(s) {
		if _, ok := counter[l]; ok {
			return false
		}
		counter[l] += 1
	}

	return true
}
