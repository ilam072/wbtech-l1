package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "snow dog sun"
	fmt.Println(reverseWords(s1))

	s2 := "hello everyone"
	fmt.Println(reverseWords(s2))

}

func reverseWords(s string) string {
	words := strings.Fields(s)
	i, j := 0, len(words)-1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}
