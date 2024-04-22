package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("I am assigned to var a")
	}
	a()
}
