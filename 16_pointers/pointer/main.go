package main

import "fmt"

func main() {
	var x int
	fmt.Println("x before:", x)
	fmt.Println("x address:", &x)
	// In go everything is passed by value
	foo(&x) // pass in the memory address of x
	fmt.Println("x after:", x)
	fmt.Println("x address:", &x)
}

// type of parameter y a place in memory that holds an int
func foo(y *int) {
	fmt.Println("y before:", y)
	fmt.Println("y value:", *y)
	// assigning a new value to *y mutates the value held at x's memory address
	*y = 43
	fmt.Println("y after:", y)
	fmt.Println("y value:", *y)
}
