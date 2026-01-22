package main

import "fmt"

func main() {
	// 1 способ: XOR-обмен
	a := 5
	b := 7

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("XOR swap (5, 7):", a, b)

	// 2 способ: сложение / вычитание
	a = 5
	b = 7

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("Add/Sub swap (5, 7):", a, b)

	// оба числа отрицательные
	a = -3
	b = -8

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("XOR swap (-3, -8):", a, b)

	// одно отрицательное число
	a = -10
	b = 4

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("Add/Sub swap (-10, 4):", a, b)

	// нулевые значения
	a = 0
	b = 0

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("XOR swap (0, 0):", a, b)
}
