package main

import "fmt"

//Function
func add(a1, a2 int) {
	res := a1 + a2
	fmt.Println("Result: ", res)
}

//Main function
func main() {

	fmt.Println("Start")

	//multiple defer statements
	//Executes in LIFO order
	defer fmt.Println("End")
	defer add(34, 46)
	defer add(10, 10)
}
