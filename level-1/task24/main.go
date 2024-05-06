package main

import (
	"fmt"
	"math"
)

//координаты точки
type Point struct {
	x float64
	y float64
}

//конструктор
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x, 
		y: y,
	}
}

//метод нахождения расстояния
func (p *Point)getDistance(other *Point) float64{
	return math.Sqrt(math.Pow(other.x - p.x, 2) + math.Pow(other.y - p.y, 2))
}

func main() {
	p1 := NewPoint(2.2, 4.4)
	p2 := NewPoint(7.3, 6.2)

	fmt.Println(p1.getDistance(p2))
}
