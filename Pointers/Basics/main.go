package main

/*
& => Address operator => pointer to get the memory address of a variable -> memory address where variable is stored

* => dereferencing operator =. get the value stored at a memory address

*int ==> pointer to an int ==> this is an address value that points to an integer elsewhere in memory
*/

import "fmt"

func main() {
	x := 42
	p := &x
	fmt.Println(x)
	fmt.Println(&x) //pointer to x (address to x)

	fmt.Println(p)  // output: memory location of x
	fmt.Println(&p) // p is stored at a dif address than x. p's value is x's address
	fmt.Println(*p) //output: 42 (value at the memory address of p, which is a pointer to the address of x. value of x = 42 => *p = 42)

	*p = 5 //changing the value found at address p (which points to x)

	fmt.Println(x)
	fmt.Println(*p)

}
