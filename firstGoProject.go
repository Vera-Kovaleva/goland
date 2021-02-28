package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello!")

	err := getRating(11)
	if err != nil {
		panic(fmt.Errorf("%w", err))
		//fmt.Println(fmt.Errorf("%w", err))
	}
}

func getRating(age int) error {
	if age < 18 {
		return errors.New("age should be 18+")
	}

	fmt.Println("Welcome on board.")

	return nil
}
