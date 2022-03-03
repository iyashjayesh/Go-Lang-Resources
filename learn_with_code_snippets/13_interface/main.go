package main

//Interfaces are named collections of method signatures.

import (
	"fmt"
	"math"
)

// Define a interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Interfaces")
	//Interfaces
	//An interface is a set of method signatures that a type must implement.
	//The interface is a contract that a type must implement.

	//Creating a struct
	c := Circle{x: 0, y: 0, radius: 5}
	r := Rectangle{width: 10, height: 5}

	fmt.Println("Circle Area:", getArea(c))
	fmt.Println("Rectangle Area:", getArea(r))

}
