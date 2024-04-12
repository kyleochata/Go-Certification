package main

import "fmt"

//creation of the person struct
type person struct {
	first string
	last  string
	age   int
}

func main() {
	//creating a value of a certain type. value of p1 is the Type of the person struct
	p1 := person{first: "James", last: "Bond", age: 22}
	//anon struct. define it then pop the val
	p2 := struct {
		first string
		last  string
		age   int
	}{"Jen", "Smith", 33}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T\t %#v\t %v\n", p1, p1, p1)
	fmt.Printf("%T\t %#v\t %v\n", p2, p2, p2)

}
