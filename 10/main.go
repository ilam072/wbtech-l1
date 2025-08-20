package main

import "fmt"

func main() {
	l := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float64)
	for _, n := range l {
		f := int(n) / 10 * 10
		m[f] = append(m[f], n)
	}

	fmt.Println(m)
}
