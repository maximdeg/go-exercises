package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 3)
	c <- "Message1"
	c <- "Message2"
	c <- "Message3"

	// Len() is the lenght of the channel, and cap() is the capacity of the channel
	fmt.Println(len(c), cap(c))

	// Range y close
	// Always close the channel
	close(c)

	// Print all the channel content with range, gives you the quantity of content
	for message := range c {
		fmt.Println(message)
	}

	// Select, is to choose which channel you want to assign to which variable
	email1 := make(chan string)
	email2 := make(chan string)
	email3 := make(chan string)
	go message("Message1", email1)
	go message("Message2", email2)
	go message("Message3", email3)
	for i := 0; i < 3; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email received email1", m1)
		case m2 := <-email2:
			fmt.Println("Email received email2", m2)
		case m3 := <-email3:
			fmt.Println("Email received email3", m3)
		}
	}
}
