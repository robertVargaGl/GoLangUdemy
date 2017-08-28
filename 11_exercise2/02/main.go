package main

import (
	"fmt"
)

func main() {
	math := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(math(2))
}
