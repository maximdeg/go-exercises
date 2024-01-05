package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	fmt.Println(f.area())
}

// Function inside the interface
func (r rectangle) area() float64 {
	return r.base * r.height
}

func (c square) area() float64 {
	return c.base * c.base
}

func main() {
	sqBase, recBase, recHeight := getInputNumbers()

	// Assigning variables to the structs
	mySquare := square{base: float64(sqBase)}
	myRectangle := rectangle{base: float64(recBase), height: float64(recHeight)}

	// Calling the function handling the interface
	fmt.Print("Square area: ")
	calculate(mySquare)
	fmt.Print("Rectangle area: ")
	calculate(myRectangle)

	// Declaring a list of interfaces

	myInterface := []interface{}{mySquare, myRectangle}
	fmt.Println(myInterface...)

}

func getInputNumbers() (float64, float64, float64) {
	var sqBase, recBase, recHeight float64
	scanner := bufio.NewScanner(os.Stdin)
	i := false
	for i != true {
		fmt.Println("Please insert the base of the square to calculate the area:")
		scanner.Scan()
		sqBase = validateFloat(strconv.ParseFloat(scanner.Text(), 64))

		fmt.Println("Please insert the base of the rectangle to calculate the area:")
		scanner.Scan()
		recBase = validateFloat(strconv.ParseFloat(scanner.Text(), 64))

		fmt.Println("Please insert the height of the rectangle to calculate the area:")
		scanner.Scan()
		recHeight = validateFloat(strconv.ParseFloat(scanner.Text(), 64))
		if sqBase != 0.0 && recBase != 0.0 && recHeight != 0.0 {
			i = true
		} else {
			fmt.Println("Please start over and enter valid numbers")
		}
	}

	return sqBase, recBase, recHeight
}

func validateFloat(input float64, err error) float64 {
	if err != nil {
		fmt.Println("Please insert a valid number!")
	} else {
		fmt.Println("Your number:", input)
	}
	return input
}
