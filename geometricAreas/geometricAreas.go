package main

import (
	"fmt"
	"math/rand"
)

func CircleArea(x float64) {
	pi := 3.14
	area := pi * x * x
	fmt.Println("The area of the circle is:", area)
}

func TrapezeArea(y, x float64) {
	var h float64 = 4.2
	area := ((x + y) * h) / 2
	fmt.Println("The area of the trapeze is:", area)
}

func RectangleArea(y, x float64) {
	area := x * y
	fmt.Println("The area of the rectangle is:", area)

}
func main() {
	var y, x float64
	y = rand.Float64()
	x = rand.Float64()
	fmt.Printf("The first number is: %v, and the second number is: %v. \n", y, x)
	RectangleArea(y, x)
	TrapezeArea(y, x)
	CircleArea(x)
}
