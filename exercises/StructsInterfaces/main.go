package main

import "fmt"

// Deining struct for triangle and square

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}
type shape interface {
	getArea() (string, float64)
}

func main() {
	// create a triangle and circle and print their properties

	triangle1 := triangle{height: 10,
		base: 5}
	square1 := square{sideLength: 10}

	// simpler way of printing directly
	// fmt.Printf("Triangle1: %+v\n", triangle1)
	// fmt.Printf("Square1: %+v \n", square1)
	// triangle1.print()
	// square1.print()

	// Print the area of the shapes without an interface
	// triAreaText, triAreaVal := triangle1.getArea()
	// fmt.Println(triAreaText, triAreaVal)

	printArea(triangle1)
	printArea(square1)
}

// Using receiver for the struct to print out details
// func (t triangle) print() {
// fmt.Printf("Triangle: \t%+v\n", t)
// }

// func (s square) print() {
// fmt.Printf("Square: \t%+v\n", s)
// }

// Receivers for functions to get the area

func (t triangle) getArea() (string, float64) {
	return "Area of a triangle:\t", 0.5 * t.base * t.height
}

func (s square) getArea() (string, float64) {
	return "Area of a square:\t", s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
