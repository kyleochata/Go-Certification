package main

import "fmt"

//defer orders as last-in-first-out
func main() {
	x := []int{1, 2, 3, 4}
	defer fmt.Println(foo(x...))
	defer fmt.Println(1)
	fmt.Println(bar())
}

func foo(n ...int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar() (int, string) {
	return 7, "Hello World"
}
