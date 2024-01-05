package main

import "fmt"

func main() {
	var message string = "Hello world"
	easyMessage := "Hello world using go :="
	fmt.Println(message)
	fmt.Println(easyMessage)
	a := 10.
	var b float64 = 3
	var c int = 10
	d := 3
	fmt.Println(c / d)
	fmt.Println(a)
	fmt.Println(b)
	x := true
	y := false
	fmt.Println(x || y)
	private()
	Public()
	printHello()
}

func private() {
	fmt.Println("This is a private function")
}

func Public() {
	fmt.Println("This is a Public function")
}

func printHello() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
