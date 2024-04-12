package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	switch {
	case x < 4 && y < 4:
		fmt.Printf("x is: %v \n y is: %v \n Both are less than 4", x, y)
	case x > 6 && y > 6:
		fmt.Printf("x is: %v \n y is: %v \n Both are greater than 6", x, y)
	case x >= 4 && x <= 6 && y != 5:
		fmt.Printf("x is: %v \n y is: %v X is greater than or equal to 4, while also being greater than or equal to 6", x, y)
	default:
		fmt.Printf("x = %v \n y = %v \n No cases matched \n", x, y)
	}
}
