package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(4))
}

func factorial(x int) (v int) {
	v = 1
	for y := x; y > 1; y-- {
		v *= y
	}
	return v
}
