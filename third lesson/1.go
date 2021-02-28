package main

import (
	"fmt"
)

func main() {
	var firstNumber float64
	var operation string
	var secondNumber float64

	fmt.Println("enter two numbers separated by a space and between them operation that you want to do")

	fmt.Scanf("%f", &firstNumber)
	fmt.Scanf("%s", &operation)
	fmt.Scanf("%f", &secondNumber)

	var result float64
	switch operation {
	case "*":
		result = firstNumber * secondNumber
		break
	case "+":
		result = firstNumber + secondNumber
		break
	case "/":
		result = firstNumber / secondNumber
		break
	case "-":
		result = firstNumber - secondNumber
		break
	}
	fmt.Println(result)
}
