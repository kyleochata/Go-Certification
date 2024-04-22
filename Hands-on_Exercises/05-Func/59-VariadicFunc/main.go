package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(foo(x...))
}

func foo(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func bar(xi ...[]int) int {
	sum := 0
	for _, v := range xi {
		for _, num := range v {
			sum += num
		}
	}
	return sum
}
