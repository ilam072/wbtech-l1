package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("5738958459328275768695942", 10)
	b := new(big.Int)
	b.SetString("121212121212121212121212", 10)

	product := new(big.Int)
	product.Mul(a, b)
	fmt.Println(product.String()) // 695631328403427365902537728611095838996876521704

	division := new(big.Rat)
	division.SetFrac(a, b)
	fmt.Println(division.FloatString(20)) // 47.34640728945827509174

	addition := new(big.Int)
	addition.Add(a, b)
	fmt.Println(addition.String()) // 5860170580540396980817154

	subtraction := new(big.Int)
	subtraction.Sub(a, b)
	fmt.Println(subtraction.String()) // 5617746338116154556574730
}
