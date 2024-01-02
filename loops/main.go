package main

import "fmt"

func main() {

	// For condicional
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
	}

	fmt.Println("")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Print(counter)
		counter++
	}

	fmt.Println("")

	// For forever, terminar con ctrl +   C
	/*counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}*/

	for i := 20; i > 0; i-- {
		fmt.Print(i)
	}
}
