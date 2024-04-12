package main

import (
	"fmt"
	"math/rand"
)

/*
use "statement statement" idiom:

	x is a rand num b/w [0, 5)
	x ==3 => print "x is 3"

run code 100x
benefit of statement statement idiom
*/
func main() {
	var i int8
	var xCount int
	for i < 100 {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is 3")
			xCount++
		}
		i++
	}
	fmt.Printf("x was 3 %v times \n", xCount)
}
