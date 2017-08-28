package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	for x := 0; x < 10; x++ {
	// 		fmt.Println(i, " - ", x)
	// 	}
	// }

	// j := 0
	// for j < 10 {
	// 	fmt.Println(j)
	// 	j++
	// }

	c := 0

	for {
		fmt.Println(c)
		if c == 9 {
			break
		}
		c++
	}

}
