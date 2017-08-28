package main

import (
	"fmt"
)

var x = 10

func main() {

	a := 43

	fmt.Println(x)
	fmt.Println(a)

	foo()
}

func foo() {
	fmt.Println(x)
	// fmt.Println(a) = error
}
