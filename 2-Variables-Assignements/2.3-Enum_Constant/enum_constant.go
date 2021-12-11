package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	a2 = iota
	a1
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a1)
}
