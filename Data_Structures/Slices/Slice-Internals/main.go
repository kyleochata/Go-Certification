package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5} //slice that has a pointer to underlying arr
	b := a                       // b stores same stuff as a; creates a pointer to a's underlying arr

	//a is a

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("===================")
	fmt.Println("Slice")
	a[0] = 7

	//when setting a[0], b will also update due to b just being a pointer to the underlying array of a. Since a's underlying array was changed, b's value updates as well
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("===================")

	x := [3]int{1, 2, 3}
	y := x
	z := &x // Pointer to x's array (var z *[3]int = &x //this will also work)

	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
	fmt.Println("===================")
	fmt.Println("Array")
	x[0] = 19

	//when changing x for Arr's; only x changes since y doesn't get the pointer automatically made like with slices
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
	fmt.Println("===================")

	//copy internals =>
	c1 := []int{0, 1, 2, 3, 4, 5}
	c2 := make([]int, 6)

	//copy(destination, source []Type) int
	copy(c2, c1)

	//c2 is a new slice with a specific underlying array. when copying c1 to c2 the values of c1 are placed into c2, but c2 does NOT have a pointer to c1. When c1 is changed, c2 will not recieve the same change due to it having a point to the c2 underlying array rather than to c1.
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("===================")

	c1[0] = 40
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("===================")

}
