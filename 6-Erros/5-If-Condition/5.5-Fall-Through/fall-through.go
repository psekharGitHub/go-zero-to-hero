package main

import (
	"fmt"
)

//fallthrough executes all the cases starting from its occurance till the
//end except the default case
func main() {

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
		fallthrough
	case i >= 20:
		fmt.Println("greater than or equal 20")
	default:
		fmt.Println("greater than 20")
	}
}
