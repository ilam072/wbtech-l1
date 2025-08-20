package main

import "fmt"

func main() {
	fmt.Println(reverse("12345"))
}

func reverse(word string) string {
	runes := []rune(word)
	i, j := 0, len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}
