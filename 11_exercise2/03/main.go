package main

import (
	"fmt"
)

func main() {
	array := []int{12, 32, 45, 23, 56, 22, 91, 35, 56}

	top := topvalue(array...)

	fmt.Println(top)
}

func topvalue(n ...int) (topv int) {
	topv = 0
	for _, i := range n {
		if topv < i {
			topv = i
		}
		// fmt.Println(i)
	}

	return
}
