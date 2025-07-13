package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) distanceTo(p2 Point) (distance float64) {
	sqrX := math.Pow((p.x - p2.x), 2)
	sqrY := math.Pow((p.y - p2.y), 2)
	distance = math.Sqrt(sqrX + sqrY)
	return distance
}

func main() {
	p1 := newPoint(1.5, 6.0)
	p2 := newPoint(6.2, 5.3)

	distance := p1.distanceTo(p2)
	fmt.Println(distance)
}
