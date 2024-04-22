package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	x := sum(xi...) //spread/unfurl the ints in the slice xi
	fmt.Println("the sum = ", x)
}
func sum(n ...int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}
