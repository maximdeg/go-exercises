package main

import "fmt"

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{
		Width:  2,
		Height: 5,
	}
	fmt.Println("Original Circle:")
	PrintShapeProperties(c)
	fmt.Println("Original ")
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func PrintShapeProperties(c Circle) {
	fmt.Printf("Area %f\n", c.Area())
	fmt.Printf("Perimeter %f\n", c.Perimeter())

}
