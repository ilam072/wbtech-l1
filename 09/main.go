package main

import "fmt"

func main() {
	nums := gen(1, 2, 3, 4, 5, 6, 7, 8)

	numsSquares := squares(nums)

	printSquares(numsSquares)
}

func gen(numbers ...int) chan int {
	res := make(chan int, len(numbers))

	go func() {
		for _, num := range numbers {
			res <- num
		}
		close(res)
	}()

	return res
}

func squares(numbers chan int) chan int {
	res := make(chan int)

	go func() {
		for num := range numbers {
			res <- num * num
		}
		close(res)
	}()

	return res
}

func printSquares(squares chan int) {
	for n := range squares {
		fmt.Println(n)
	}
}
