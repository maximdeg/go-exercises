package main

import (
	"fmt"
	"sync"
	"time"
)

// Takes a string and a pointer to a Wait Group for synchronization
func say(text string, wg *sync.WaitGroup) {
	// Uses defer to ensure that wg is called even if the function exits early
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	// Adds 1 indicating one more goroutine to wait for
	wg.Add(2)

	go say("World", &wg)
	go say("Forever", &wg)

	// Blocks the main program until the counter reaches zero
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	// Pauses the program for 2 sec, allowing the goroutines to finish before the program does
	time.Sleep(time.Second * 1)
}
