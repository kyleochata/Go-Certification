package main

import "fmt"

/*
● start with this slice
	○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
● append to that slice this value
	○ 52
● print out the slice
● in ONE STATEMENT append to that slice these values
	○ 53
	○ 54
	○ 55
● print out the slice
● append to the slice this slice
	○ y := []int{56, 57, 58, 59, 60}
● print out the slic
*/
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	fmt.Println("=======================")

	x = append(x, 53, 54, 55)
	fmt.Println(x)
	fmt.Println("=======================")
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) // when appending a slice to a slice, the slice that's getting added needs to be spread
	x = append(x, []int{1, 2, 3}...)
	fmt.Println(x)
}
