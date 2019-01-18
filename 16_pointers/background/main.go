package main

import "fmt"

func main() {
	a := 42 // assign a value to a variable
	fmt.Println(a)
	b := &a // b is now equal to the memory address of a
	fmt.Println(b)

	fmt.Printf("%T\n", a) // type of a is still an int
	fmt.Printf("%T\n", b) // type of b is pointer to an int (*int)

	fmt.Println(*b)  // prints the value stored at memory address b
	fmt.Println(*&a) // get the value stored at the memory address of a
	fmt.Println(*b)  // same as line above

	*b = 43        // change the value stored at b
	fmt.Println(a) // a is now 43 because it pointing to the value stored at b
}
