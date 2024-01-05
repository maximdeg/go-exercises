package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)  // Prints the value of x
	fmt.Println(&x) // Prints the memory address of cx
	changeValue(x)
	y := &x
	fmt.Println(y)  // Now this is the memory address of x
	fmt.Println(*y) // And this prints the value of x because is pointing to the memory address
}

func changeValue(a int) {
	a = 36
	fmt.Println(&a, a) // Changes the value of the memory address and variable and value
}
