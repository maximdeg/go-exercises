package main

import "fmt"

// Instance of the interface
type animal interface {
	move() string
}

// Structs definitions
type dog struct{}
type fish struct{}
type bird struct{}

// All the move functions
func (dog) move() string {
	return "I'm a dog and I walk"
}

func (fish) move() string {
	return "I'm a fish and I swim"
}

func (bird) move() string {
	return "I'm a bird and I fly"
}

// Function using the interface
func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	d := dog{}
	f := fish{}
	b := bird{}
	moveAnimal(d)
	moveAnimal(f)
	moveAnimal(b)
}
