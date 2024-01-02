package main

import (
	"fmt"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Is 1")
	} else {
		fmt.Println("Is not 1")
	}

	// with AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("AND is True")
	}

	// with OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Is true, OR")
	}

	/*value3, err := strconv.Atoi("asdasd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value3)*/

	// Switch

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Is even")
	default:
		fmt.Println("Is odd")
	}

	// Switch without condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("Is more than 100")
	case value < 0:
		fmt.Println("Is less than 0")
	default:
		fmt.Println("Is either, is:", value)
	}
}
