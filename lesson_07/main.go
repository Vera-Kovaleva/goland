package main

import "fmt"

type animal interface {
	walker
	runner
}

type walker interface {
	walk()
}
type runner interface {
	run()
}

func action(a animal) {
	a.walk()
	a.run()
}

type cat struct{}

func (c *cat) walk() {
	fmt.Println("cat is walking")
}
func (c *cat) run() {
	fmt.Println("cat is running")
}

type dog struct{}

func (c *dog) walk() {
	fmt.Println("dog is walking")
}
func (c *dog) run() {
	fmt.Println("dog is running")
}

func main() {
	var c = &cat{}
	var d = &dog{}

	action(c)
	action(d)
}
