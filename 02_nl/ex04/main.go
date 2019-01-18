package main

import "fmt"

func main() {
	x := 42
	log(x)
	y := x << 1
	log(y)
}

func log(x int) {
	fmt.Printf("%d\t%b\t%#X\n", x, x, x)
}
