package main

import (
	"fmt"
)

func main() {
	var i int
	var number int
	var numberI int
	fmt.Println("enter a value")
	fmt.Scanf("%d", &number)
	for i = 0; i < 3; i++ {
		numberI = number % 10
		fmt.Println(numberI)
		number = number / 10
	}

}
