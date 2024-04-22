package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
	y := bar()
	fmt.Println(y())

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", bar)

}

func foo() int {
	return 4
}

// bar is a func that returns a func that returns an int
func bar() func() int {
	return func() int {
		return 48
	}
}
