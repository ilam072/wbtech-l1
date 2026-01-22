package main

import "fmt"

func main() {
	fmt.Printf("%04b\n", 5)
	fmt.Printf("%04b\n\n", setBit(5, 1))

	fmt.Printf("%04b\n", 5)
	fmt.Printf("%04b\n", resetBit(5, 1))
}

func setBit(num int64, i int) int64 {
	if i < 0 || i > 63 {
		return num
	}

	var mask int64 = 1 << (i - 1)
	return num | mask
}

func resetBit(num int64, i int) int64 {
	if i < 0 || i > 63 {
		return num
	}

	var mask int64 = 1 << (i - 1)
	return num &^ mask
}
