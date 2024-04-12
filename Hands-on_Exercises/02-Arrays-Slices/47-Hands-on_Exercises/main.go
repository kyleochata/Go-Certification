package main

import "fmt"

/*
Create a slice to store the names of all of the states in the United States of America.
	○ Use make and append to do this.
	○ Goal: do not have the array that underlies the slice created more than once.
● Print out
	○ the len
	○ the cap
	○ the values, along with their index position, without using the range clause.
*/
func main() {
	xs := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `Kentucky`, `Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `cUtah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Printf("Length: %v\t capacity: %v\n", len(xs), cap(xs))
	for i := 0; i < len(xs); i++ {
		fmt.Printf("index: %v\t State Name: %s\n", i, xs[i])
	}
}
