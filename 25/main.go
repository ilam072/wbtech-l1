package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Start:", time.Now())
	sleep(3 * time.Second)
	fmt.Println("End:", time.Now())
}
