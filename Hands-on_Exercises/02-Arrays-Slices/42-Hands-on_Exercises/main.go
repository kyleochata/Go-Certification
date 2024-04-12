package main

import "fmt"

/*
Composite literal ==> create an arr to hold 5 values of int
assign values to ea index pos
Range over array and print out type and values
*/
func main() {
	a := [5]int{0, 1, 5, 3, 4}
	for i, v := range a {
		fmt.Printf("position: %v\t value: %v\n", i, v)
	}
}
