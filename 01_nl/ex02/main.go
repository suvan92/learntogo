package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	fmt.Printf("%v, %v, %v\n", x, y, z)

	s := fmt.Sprintf("%v - %v - %v", x, y, z)
	fmt.Println(s)
}
