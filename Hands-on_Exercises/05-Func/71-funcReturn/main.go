package main

import "fmt"

func main() {
	fmt.Println(apple()()) // Call the apple() function before passing it as an argument.
	//double () => apple() is called and returns anon fxn ==> second () => returns the string
	//apple() Output: 0xb13c20
}

func apple() func() string {
	return func() string {
		a := "apples"
		return a
	}
}
