package main

import "fmt"

//JSON payload are string so big in size
//'proto' uses byte array as payload which are smaller and hence faster than JSON
//new form of payload data called 'proto', instead of JSON
func main() {
	str := "MBB$"
	data := []byte(str)
	fmt.Println(data)
}
