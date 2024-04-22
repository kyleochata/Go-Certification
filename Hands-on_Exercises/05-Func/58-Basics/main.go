package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 3
}

func bar() (int, string) {
	return 7, "Hello World"
}
