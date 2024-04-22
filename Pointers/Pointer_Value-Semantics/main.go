package main

import "fmt"

//Passing in a value and return a value => value semantics
func addOne(v int) int {
	return v + 1
}

//Pointer Semantics
func addOneP(v *int) {
	*v += 1
}

func main() {
	//Value semantics
	a := 1
	fmt.Println(a)         //1
	fmt.Println(addOne(1)) //2
	fmt.Print(a)           //1

	b := 1
	fmt.Println("========")
	fmt.Println(b)
	addOneP(&b)
	addOneP(&a)
	fmt.Println(b)
	fmt.Println(a, "aaaaaaaa")
}
