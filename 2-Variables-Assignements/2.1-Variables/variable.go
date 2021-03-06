package main

import "fmt"

func main() {
	//var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	//You can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println(b, c)

	//Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)
}
