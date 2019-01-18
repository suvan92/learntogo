package main

import "fmt"

const (
	year = 2017
	a    = year + iota
	b    = year + iota
	c    = year + iota
	d    = year + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
