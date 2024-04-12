package main

import "fmt"

/*
Composite literal ==> create a slice of int => assign 10 values.
Range over array and print out type and values
*/
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 19}
	for i, v := range a {
		fmt.Printf("position: %v\t value: %v\t type: %T\n", i, v, v)
	}
}
