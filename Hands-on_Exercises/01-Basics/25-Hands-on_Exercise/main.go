package main

import (
	"fmt"
)

func main() {
	var x int
	for x < 50 {
		if x%3 == 0 {
			fmt.Printf("%v = x\n", x)
		}
		x++
	}
}
