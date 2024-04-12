package main

import "fmt"

func main() {
	{
		//declare a avariable of type [7]int ==> array with 7 ints in it
		var ni [7]int

		//assign value to ni[0]
		ni[0] = 40
		fmt.Printf("%v \t\t\t\t %T\n", ni, ni)

		//array literal
		ni2 := [4]int{55, 54, 56, 59}
		fmt.Printf("%v \t\t\t\t %T\n", ni2, ni2)

		//array literal ==> have the compiler count elements
		//it [...] ==> have compiler count how many items should be in the array you are creating
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		//use built-in len fxn
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}
}
