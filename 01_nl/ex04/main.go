package main

import "fmt"

type chelsea int

var x chelsea

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
}
