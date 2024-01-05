package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secret := getRandomNumber()
	for matching := false; !matching; {
		guess := getInputNumber()

		//fmt.Println(secret, guess)
		matching := isMatching(secret, guess)
		if matching {
			break
		}
	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your number is too big")
		return false
	} else if guess < secret {
		fmt.Println("Your number is too small")
		return false
	} else {
		fmt.Println("CONGRATULATIONS!!")
		return true
	}
}

func getInputNumber() int {
	fmt.Println("Please insert a number between 1 and 10:")
	var input int

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Please insert a valid number!")
	} else {
		fmt.Println("Your number:", input)
	}
	return input
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
