package main

import (
	"fmt"
)

func main() {
	t(1)
	t("hello!")
	t(true)

	chi := make(chan int)
	t(chi)

	chs := make(chan string)
	t(chs)

	chb := make(chan bool)
	t(chb)

	t(struct{}{})
}

func t(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v type of int\n", v)
	case string:
		fmt.Printf("%v type of string\n", v)
	case bool:
		fmt.Printf("%v type of bool\n", v)
	case chan int:
		fmt.Printf("%v type of chan int\n", v)
	case chan string:
		fmt.Printf("%v type of chan string\n", v)
	case chan bool:
		fmt.Printf("%v type of chan bool\n", v)
	default:
		fmt.Printf("unexpected type %T!\n", v)
	}
}
