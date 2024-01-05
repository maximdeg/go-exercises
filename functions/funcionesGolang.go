package main

import "fmt"

// Funcion normal sin return
func normalFunction(message string) {
	fmt.Println(message)
}

// Funcion normal recibiendo multiples variables
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

// Funcion devolviendo un int y recibiendo un int
func returnValue(a int) int {
	return a * 2
}

// Funcion recibiendo un int y devolviendo 2 variables int
func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	// Declarando 2 variables a una funcion que devuelve 2 valores
	value1, value2 := doubleReturn(3)
	fmt.Println("Value1= ", value1, "Value2= ", value2)

	// Declarando 1 variable a una funcion que devuelve 1 valor
	value1, _ = doubleReturn(4)
	fmt.Println("Value1= ", value1)
}
