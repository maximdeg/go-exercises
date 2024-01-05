package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Raw strings to check the operators that the calculator will use
const (
	// Checks if the user will insert a float
	operationRegexp string = `[0-9]+(\.[0-9]+)?[\+\-\*/][0-9]+(\.[0-9]+)?`
	// Checks the operator of the operation
	operatorsRegexp string = `[\+\-\*/]`
)

func getInput() string {
	fmt.Println("Insert an operation:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func obtainOperation() (string, error) {
	operation := getInput()
	expr := operationRegexp
	patern := regexp.MustCompile(expr)
	// We check if the operation is really and operation, otherwise return error
	if patern.MatchString(operation) {
		return operation, nil
	}
	return "", errors.New("Lexical: Input is not an operation")
}

func obtainNumbers(operation string) (float64, float64) {
	expr := operatorsRegexp
	patern := regexp.MustCompile(expr)
	numbers := patern.Split(operation, 2)
	operand1, _ := strconv.ParseFloat(numbers[0], 32)
	operand2, _ := strconv.ParseFloat(numbers[1], 32)
	return operand1, operand2
}

func obtainOperator(operation string) byte {
	expr := operatorsRegexp
	patern := regexp.MustCompile(expr)
	indexes := patern.FindIndex([]byte(operation))[0]
	return operation[indexes]
}

// Gets the operation on a string and breaks it down
func operate() float64 {
	operation, err := obtainOperation()
	if err != nil {
		os.Exit(1)
	}
	value1, value2 := obtainNumbers(operation)
	return calculate(value1, value2, obtainOperator(operation))
}

func calculate(value1 float64, value2 float64, operator byte) float64 {
	var ans float64
	switch operator {
	case '+':
		ans = value1 + value2
	case '-':
		ans = value1 - value2
	case '*':
		ans = value1 * value2
	case '/':
		ans = value1 / value2
	}
	return ans
}

func main() {
	ans := operate()
	fmt.Printf("The output value is %f\n", ans)
}
