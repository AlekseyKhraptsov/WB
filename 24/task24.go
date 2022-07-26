package main

// Задание 24
// Разработать программу нахождения расстояния между двумя точками, которые представлены
// в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// NewPoint создаёт новую точку.
func newPoint(x, y float64) Point {
	return Point{x, y}
}

// расчет растояния
func (p Point) Distance(q Point) float64 {
	square := func(z float64) float64 { return z * z }
	return math.Sqrt(square(q.x-p.x) + square(q.y-p.y))
}

func main() {
	p1 := newPoint(5, -2)
	p2 := newPoint(3, 4)
	fmt.Printf("Distance: %.2f\n", p1.Distance(p2))

}
