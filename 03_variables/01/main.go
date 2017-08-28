package main

import (
	"fmt"
)


// https://godoc.org/fmt

func main() {
	fmt.Println("------------------------------")


	// declatare variables
	a := 13
	b := false



	// general printing
	fmt.Printf("%v - %%v - value of the default format \n", a)
	fmt.Printf("%T - %%T - print type of variabe \n", a)
	fmt.Println("")
	
	
	// boolean
	fmt.Printf("%t - %%t - obly boolean variable \n", b)
	fmt.Println("")
	

	// fmt.Printf("%v \n", a)
	// fmt.Printf("%T \n", a)
	// fmt.Printf("%t \n", d)
	fmt.Println("------------------------------")
}