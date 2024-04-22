package main

import (
	"fmt"
)

//intDelta dereferences pointer to int to reassign value
/*
1. a is a var with Type: int and Val: 42
2. intDelta takes a pointer to an int
	a) we feed a pointer to a to the intDelta
	b) &a gets reassigned to var n *int
	c) dereference n (which is &a) to get the value found at &a
*/
func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}
func mapDelta(ii map[string]int, s string) {
	ii[s] = 99
}

func main() {
	a := 42
	intDelta(&a) //we need to feed intDelta &a because a is just a Type of int. &a is Type of *int(pointer)
	fmt.Println(a)

	//slices include a pointer to the underlying array. Don't need to & or * to mutate values
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	//maps also have a pointer to its underlying collect ion of key-value pairs.
	m := make(map[string]int)
	m["James"] = 54
	fmt.Println(m)
	mapDelta(m, "James")
	fmt.Println(m)
}

/*
Slices, maps, pointers are all mutable data in Go (can be changed)
	even though they are passed by value
		behave as if they were passed by ref => "value" is copied and passed = reference to underlying data, NOT THE ACTUAL DATA
ex:
````
nums := []int{1,2,3}

func modify (s []int){
	s[0] = 100
}
modify(nums) // output: [100, 2, 3]
````

Slice header includes pointer to the underlying arr is passed by value
	Copied value (pointer) is still pointing to same underlying arr ==> changes made inside the fxn are visible outside of it even though the data was passed by value
	Dereferencing the pointer and modifying the value it points to inside of a function will modify the original value of the address where the pointer points to

In Go, mutable data types  = "value" that's passed includes a ref to the underlying data
	changes made inside the function are visible outside of it

*/
