package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

type Square struct {
	side float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (s Square) area() float64 {
	return s.side * s.side
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 5, height: 10}
	square := Square{side: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
	fmt.Printf("Square area: %f\n", getArea(square))
}
