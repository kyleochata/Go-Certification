package main

import "fmt"

func main() {
	v := sum(1, 2, 3, 4, 5)
	fmt.Println(v)
}

//by using the variadic param (...int), the func sum will take all those ints and create a slice out of it ([]int)
func sum(i ...int) int {
	fmt.Println(i)
	fmt.Printf("%T\n", i) // []int

	//because GO made a slice automatically for us, we can use a for range to add all ints up
	var sum int
	for _, v := range i {
		sum += v
	}
	return sum
}
