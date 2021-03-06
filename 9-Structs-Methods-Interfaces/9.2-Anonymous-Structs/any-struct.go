package main

import "fmt"

type Kitchen struct {
	numOfPlates int
	numOfSpoons int
}

type House struct {
	Kitchen    //anonymous field
	numOfRooms int
}

func main() {
	h := House{Kitchen{10, 5}, 3}                                     //to initialize you have to use composed type name.
	fmt.Println("House h has this many rooms:", h.numOfRooms)         //numOfRooms is a field of House
	fmt.Println("House h has this many plates:", h.numOfPlates)       //numOfPlates is a field of anonymous field Kitchen, so it can be referred to like a field of House
	fmt.Println("House h has this many spoons:", h.numOfSpoons)       //numOfSpoons is a field of anonymous field Kitchen, so it can be referred to like a field of House
	fmt.Println("The Kitchen contents of this house are:", h.Kitchen) //we can refer to the embedded struct in its entirety by referring to the name of the struct type
}
