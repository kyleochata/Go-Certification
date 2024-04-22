package main

import "fmt"

func main() {
	f := Increase()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

//to document code (go doc -cmd (-all or -u)) start with the identifier for the Value
//Increase is closing around the anon function that increases and returns x
func Increase() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
