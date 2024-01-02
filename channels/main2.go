package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a channel of integers
	numbers := make(chan int)

	// Does not reapeat sequences
	rand.Seed(time.Now().UnixNano())

	// Launch a goroutine to generate random numbers
	go func() {
		for i := 0; i < 10; i++ {
			numbers <- rand.Intn(100)
			time.Sleep(1 * time.Second) // Simulate work
		}
		close(numbers) // Close the channel when done
	}()

	// Main program receives and prints the numbers
	for number := range numbers {
		fmt.Println("Received:", number)
	}
}
