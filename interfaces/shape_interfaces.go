package main

import (
	"fmt"
	"math"
)

// Geometry interface
type Geometry interface {
	area() float64
	perim() float64
}

// Rectangle struct that hold the data
type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

// Circle struct that hold the data
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Printf("Measure called for %T: ", g)
	fmt.Println(" Area: ", g.area())
	fmt.Println(" Reactangle: ", g.perim())
}

func main() {
	rect := Rectangle{width: 3, height: 4}
	measure(rect)

	circ := Circle{radius: 5}
	measure(circ)
}
