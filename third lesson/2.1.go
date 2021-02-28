package main

import (
	"fmt"
)

func isPrimeNaive(number uint) bool {
	if number <= 1 {
		return false
	}

	var i uint
	for i = 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Print("enter number: ")

	var number uint
	fmt.Scanf("%d", &number)

	var i uint
	for i = 1; i < number; i++ {
		isCurPrime := isPrimeNaive(i)
		if isCurPrime {
			fmt.Printf("Found prime number %d\n", i)
		}
	}
}
