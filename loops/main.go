package main

import "fmt"

func main() {

	// For condicional from 0 to 10
	fmt.Println("For with conditional declaring variable:")
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
	}

	fmt.Println("")

	//For while counter equals 10
	fmt.Println("For with an outside variable, while counter equals 10:")
	counter := 0
	for counter <= 10 {
		fmt.Print(counter)
		counter++
	}

	fmt.Println("")

	// For forever, end with ctrl +   C
	/*counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}*/

	//For backwards from 20 until 0
	fmt.Println("For backwards using conditional:")
	for i := 20; i >= 0; i-- {
		fmt.Print(i)
	}
	fmt.Println("")
}
