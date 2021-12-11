package main

import "fmt"

//point is an x,y coordinate
type point struct {
	x, y int
}

func main() {
	//Go provides several 'verbs' to format general Go Values
	p := point{1, 2}
	fmt.Println("%v\n", p)

	//%+v will include struct field names
	fmt.Printf("%+v\n", p)

	//prints Go syntax representative of a value
	fmt.Printf("%#v\n", p)

	//print the type of a value
	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)
}

//%v - values
