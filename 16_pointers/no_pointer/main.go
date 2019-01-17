package main

import "fmt"

func main() {
	var x int
	fmt.Println("x before:", x)
	fmt.Println("x address:", &x)
	// In go everything is passed by value
	foo(x)
	fmt.Println("x after:", x)
	fmt.Println("x address:", &x)
}

// a new var 'y' is created and the value is the value of the argument
func foo(y int) {
	fmt.Println("y before:", y)
	fmt.Println("y address:", &y)
	// modifying y here is changing the value at a new memory address (of y)
	y = 43
	fmt.Println("y after:", y)
	fmt.Println("y address:", &y)
}
