package main

import "fmt"

/*
Recursion: func that calls itself until a base case is reached (to avoid infinite recursion)
*/

func main() {
	fmt.Println(factorial(3))
	fmt.Println(factLoop(3))
}
func factorial(n int) int {
	fmt.Println(n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//anything that can be done with recursion can be filled by a loop
func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
