package main

import "fmt"

/*
go run -gcflags -m // shows escape analysis
*/

func main() {
	a := 1
	b := 1
	fmt.Println(a)
	fmt.Println("========")
	fmt.Println(&b)
}
