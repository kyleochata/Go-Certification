package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("my name is", p.first, p.age, "years old")
}
func main() {
	p1 := person{"Kyle", 20}
	p1.speak()
}
