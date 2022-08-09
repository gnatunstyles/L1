package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func count(fst, scd *Point) float64 {
	return math.Sqrt(math.Pow((scd.x-fst.x), 2) + math.Pow((scd.y-fst.y), 2))
}

func main() {
	first := New(12.0, 32.0)
	second := New(54.0, 0.0)
	fmt.Println(count(first, second))
}
