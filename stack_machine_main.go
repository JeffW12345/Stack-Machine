package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processInput(inputAsString string) (int, error) {
	fmt.Println("Input: " + inputAsString)

	inputAsSlice := strings.Split(inputAsString, " ")

	stack := make([]int, 0)

	for _, v := range inputAsSlice {
		var err error
		switch v {
		
		case "POP":
			stack, err = popActions(stack)
		case "DUP":
			stack = dupActions(stack)
		case "+":
			stack, err = plusSymbolActions(stack)
		case "-":
			stack, err = minusSymbolActions(stack)
		case "*":
			stack, err = multiplySymbolActions(stack)
		case "CLEAR":
			stack = make([]int, 0)
		case "SUM":
			stack, err = sumActions(stack)

		default:
			numberAsInt, err := strconv.Atoi(v)
			if err != nil {
				return 0, fmt.Errorf("invalid input %s: %v", v, err)
			}
			if numberAsInt < 0 || numberAsInt > 50000 {
				return 0, fmt.Errorf("number input %s needs to be between 0 and 50,000 inclusive", v)
			}
			stack = append(stack, numberAsInt)
		}
		if err != nil{
			return 0, err
		}
	}

	if len(stack) > 1 {
		return 0, fmt.Errorf("there should be just one element in the stack at the end of the operation but there are %d elements", len(stack))
	}

	return stack[0], nil
}

func main() {
	// Example usage

	result, err := processInput("99")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	fmt.Printf("Result: %d\n", result)
}
