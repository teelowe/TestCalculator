package main

import (
	"TestCalculator/calculator"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatal("Usage: main <operand1> <operator> <operand2>")
	}

	operand1 := parseOperand(args[0])
	operand2 := parseOperand(args[2])
	operator := args[1]

	result, err := calculator.Calculate(operand1, operand2, operator)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %s %s = %f\n", args[0], args[1], args[2], result)
}

func parseOperand(operand string) float64 {
	value, err := calculator.ParseOperand(operand)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
