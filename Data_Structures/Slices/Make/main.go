package main

import "fmt"

func main() {
	//slice literal aka composite literal
	si := []string{"a", "b", "c"}
	fmt.Println(si)

	//array literal
	as := [3]string{"x", "y", "z"}
	fmt.Println(as)

	//make([]Type, size ...Type) Type
	//first arg = slice initializer with type
	//second arg =
	//third arg = how many elements are going to be found in the new slice = len(xi)
	/*
		make will create a new slice ontop of an array. In our example we are creating slice that is empty (make([]int, 0)), the 10 is signaling that even though the slice is currently empty, I want the underlying array to have a capacity of an int
	*/
	//xi := make([]int, 4, 10) //output: [0, 0, 0, 0]
	xi := make([]int, 0, 10)

	fmt.Println(xi) //output: []

	fmt.Println(len(xi)) //0

	//cap() find the capacity of the underlying array; int that goes in the brackets for arr literal
	fmt.Println(cap(xi))

	xi = append(xi, 10, 2, 4, 8, 99)
	fmt.Println(xi)

}
