package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func main() {
	p1 := NewPoint(3, 5)
	p2 := NewPoint(7, 9)

	fmt.Println(p1.Distance(p2))
}

func (p *Point) Distance(other Point) float64 {
	res := math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
	return res
}
