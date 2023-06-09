package main

import "fmt"

type OneRectangle struct {
	x float64
	y float64
}

type car struct {
	brand string
	model string
	year  int
}

func (r OneRectangle) area() float64 {
	return r.x * r.y
}

func (r OneRectangle) perimeter() float64 {
	return 2 * (r.x + r.y)
}

func (c car) presentation() {
	fmt.Printf("Hello, I am an %s model %s car %d.", c.model, c.brand, c.year)
}

func main() {
	r := OneRectangle{4, 8}
	fmt.Println("Area : ", r.area())
	fmt.Println("Perimeter :", r.perimeter())

	p := &r
	p.x = 1e9
	fmt.Println(p)

	c := car{"Volvo", "XC40 T3 FWD Momentum", 2020}
	c.presentation()
}
