package main

import "fmt"

/*
Interfaces in Go define a set of method signatures
Polymorphism = ability of a VALUE of a certain TYPE to also be of another TYPE
	In GO: VALUEs can be of more than one type
*/

type person struct {
	first string
}
type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}
func (sa secretAgent) speak() {
	fmt.Println("I am a secret", sa.first)
}

//any VALUE that has the method Type of speak() ==> human interface
//interfaces ==> if you have the right methods, then you're that interface's Type
type human interface {
	speak()
}

//can take any Value with the Type of human ==> you need to have the speak() method
func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	p2 := person{
		first: "Jen",
	}
	//call the method via dot.notation
	sa1.speak()
	p2.speak()
	//if you are of type human, use the speak method in saySomething
	//sa1 and p2 satisfy human and their own structs! == Polymorphism (Values are of one or more Types!)
	saySomething(sa1)
	saySomething(p2)
}
