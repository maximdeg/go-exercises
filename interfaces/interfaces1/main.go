package main

import "fmt"

// Create the interface
type figures interface {
	area() float64
}

// Define the structs
type rectangle struct {
	base   float64
	height float64
}

type square struct {
	base float64
}

// Function calling the interface
func calculate(f figures) {
	fmt.Println("Area: ", f.area())
}

// Function inside the interface
func (r rectangle) area() float64 {
	return r.base * r.height
}

func (c square) area() float64 {
	return c.base * c.base
}

func main() {
	// Assigning variables to the structs
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}

	// Calling the function handleing the interface
	calculate(mySquare)
	calculate(myRectangle)

	// Declaring a list of interfaces
	myInterface := []interface{}{mySquare, myRectangle, 4.90}
	fmt.Println(myInterface...)
}
