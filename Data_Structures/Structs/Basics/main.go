package main

import "fmt"

//creation of the person struct
type person struct {
	first string
	last  string
	age   int
}

func main() {
	//composite literal to make a struct with the Type of person (must fulfill all the requirements of the person struct (have first, last string and age int) before being created)
	p1 := person{first: "James", last: "Bond", age: 22}
	p2 := person{"Jen", "Smith", 33}
	//Syntax for p1 and p2 work the exact same. No errors in compiler or terminal

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T\t %#v\t %v\n", p1, p1, p1)
	fmt.Printf("%T\t %#v\t %v\n", p2, p2, p2)

}
