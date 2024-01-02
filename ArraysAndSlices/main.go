package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println("Array: ", array, len(array), cap(array))

	// Slice

	// Basic way to declare a slice
	// slice := []int{0, 1, 2, 3, 4, 5, 6}

	// Makes sure that when you use random makes always different sequences
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, rand.Intn(6)+5) // Create the slice of random amount of slots btw 5-17
	for i := 0; i < len(slice); i++ {    // Fill the slice with random numbers from 0-19
		slice[i] = rand.Intn(20)
	}

	fmt.Println("Slice: ", slice, len(slice), cap(slice))

	// Slice methods (slicing)
	fmt.Println("Slice's first number: ", slice[0])               // Only index 0
	fmt.Println("Slice, first 3 numbers", slice[:3])              // From de beginnig of the slice until 3
	fmt.Println("Slice, Slice, 3rd and 4th number: ", slice[2:4]) // From the first number(inclusive) until the second(exclusive)
	fmt.Println("Slice, from 4th number: ", slice[4:])            // From the first until the last one

	// Append: Make the slice the size you want. Ex 0-5
	if len(slice) > 5 {
		slice = append(slice, 5)
		fmt.Println(slice)
	} else {
		// Append new slice to the end of the slice
		newSlice := []int{8, 9, 10}
		slice = append(slice, newSlice...)
		fmt.Println(slice)
	}
}
