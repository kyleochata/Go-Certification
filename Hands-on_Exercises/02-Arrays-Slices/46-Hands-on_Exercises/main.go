package main

import "fmt"

/*
● start with this slice
○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
● use APPEND & SLICING to get these values here which you should ASSIGN to the
variable “x” and then print:
○ [42, 43, 44, 48, 49, 50, 51]
*/
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
	fmt.Println("=======================")
}
