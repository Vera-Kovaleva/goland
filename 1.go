package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("enter the value of two strings separated by a space")
	fmt.Scanf("%d", &a)
	var b int
	fmt.Scanf("%d", &b)
	var ploshad int = a * b
	fmt.Println(ploshad)
}
