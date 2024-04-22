package main

import (
	"fmt"
	"math"
)

func main() {
	x := pow(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func pow(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
