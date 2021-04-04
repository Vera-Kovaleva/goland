package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint32 = 1
	fmt.Println(unsafe.Sizeof(a))
	var b = &a
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := *b
	fmt.Println("c=", c)
}
