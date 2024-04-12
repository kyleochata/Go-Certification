package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("%v\n", i)
	// }
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("iteration number: %v \t number: %v \t type: %T \t binary: %b\n", i, x, x, x)

		case 1:
			fmt.Printf("iteration number: %v \t number: %v \t type: %T \t binary: %b\n", i, x, x, x)

		case 2:
			fmt.Printf("iteration number: %v \t number: %v \t type: %T \t binary: %b\n", i, x, x, x)

		default:
			fmt.Printf("iteration number: %v \t number: %v \t binary: %b\t Failed 3 cases\n", i, x, x)
		}
	}

}
