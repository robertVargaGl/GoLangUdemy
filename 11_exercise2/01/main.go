package main

import (
	"fmt"
)

func main() {
	fmt.Println(math(1))
	fmt.Println(math(2))
}

func math(number int) (int, bool) {
	ret1 := number / 2

	var ret2 bool

	if number%2 == 1 {
		ret2 = false
	} else {
		ret2 = true
	}

	return ret1, ret2
}
