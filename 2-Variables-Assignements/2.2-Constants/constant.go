package main

import (
	"fmt"
	"math"
)

//const declares a constant value
const s string = "constant"

func main() {
	fmt.Println(s)
	//A const statement can appear anywhere a var statement can
	const n = 500000000
	//n = 0
	//Constant expression perform arithmetic with arbitrary decision
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Printf("%T", d)
	//A numeric constant has no ttype untill it's given one such as by an explicit cast
	fmt.Println(int64(d))
	fmt.Printf("%T", d)

	//A number can be given a type by using it in a context that requires one, such as a variable assignment
	fmt.Println(math.Sin(n))
}
