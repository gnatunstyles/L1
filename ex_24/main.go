package main

import (
	"fmt"
	"math"
)

//структура точки
type Point struct {
	x float64
	y float64
}

//конструктор структуры точки
func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

//функция вычисления расстояния
func count(fst, scd *Point) float64 {
	return math.Sqrt(math.Pow((scd.x-fst.x), 2) + math.Pow((scd.y-fst.y), 2))
}

func main() {
	first := New(12.0, 32.0)
	second := New(54.0, 0.0)
	fmt.Println(count(first, second))
}
