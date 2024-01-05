package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	// Create the channel
	chanel := make(chan string)
	// Making a slice of servers
	servers := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0
	for {
		if i > 2 {
			break
		}
		for _, server := range servers {
			go checkServer(server, chanel) // Goroutine
		}
		time.Sleep(1 * time.Second) // Makes the program sleep for a second
		fmt.Println(<-chanel)
		i++
	}

	timePassed := time.Since(start)
	fmt.Printf("Execution time %s\n", timePassed)
}

func checkServer(server string, chanel chan string) {
	_, err := http.Get(server)
	if err != nil {
		chanel <- server + " Is no available=("
	} else {
		chanel <- server + " Is working fantastic =P" // Calling the channel
	}
}
