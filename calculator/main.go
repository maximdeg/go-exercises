package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Scan the answer
	fmt.Println("Insert operation:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()

	re := regexp.MustCompile("[^0-9]+")
	operatorr := re.FindAllString(operation, -1)
	//fmt.Println(operatorr)

	var operator string = strings.Join(operatorr, "")
	values := strings.Split(operation, operator)
	num1, _ := strconv.Atoi(values[0])
	num2, _ := strconv.Atoi(values[1])

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Operation not valid")
	}
}
