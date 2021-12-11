package main

import "fmt"

func main() {
	//Here we use range to sum the numbers in a slice
	nums := []int{2, 3, 4}
	sum := 0

	for i, num := range nums {
		sum += num
		fmt.Println("i:", i)
	}
	fmt.Println("sum: ", sum)
	//range on arrays and slice provides both index and value for each element
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}
}
