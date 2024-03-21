package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func CalculateDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func main() {
	a := NewPoint(1, 2)
	b := NewPoint(1, 3)
	fmt.Println(CalculateDistance(a, b))
}

/*
Разработать программу нахождения расстояния между двумя точками,которые представлены
в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
