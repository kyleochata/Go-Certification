package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I am", p.first) //eq to "this" (JS)
}
func main() {
	p1 := person{"Kyle"}
	p2 := person{"Jen"}

	p1.speak()
	p2.speak()

}

//methods are added to Types via the (r receiver) in func (r receiver)identifier(p params)(return(s)) { <code> }
