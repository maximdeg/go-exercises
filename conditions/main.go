package main

import (
	"fmt"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// with AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("AND es verdad")
	}

	// with OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	/*value3, err := strconv.Atoi("asdasd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value3)*/

	// Switch

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Switch without condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No es ninguna, es:", value)
	}
}
