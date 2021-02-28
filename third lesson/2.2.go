package main

import "fmt"

func main() {

	fmt.Print("enter number: ")

	var number int
	fmt.Scanf("%d", &number)

	for i := 2; i <= number; i++ {

		isPrime := true

		for j := 2; j < i; j++ {

			if i%j == 0 {

				isPrime = false
			}
		}

		if isPrime == true {

			fmt.Println(i)
		}
	}
}
