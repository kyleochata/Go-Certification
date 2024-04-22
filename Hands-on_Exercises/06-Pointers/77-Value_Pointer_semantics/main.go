package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{"June"}
	p2 := &person{"Mary"}

	fmt.Println(p1, p2)

	fmt.Println(changeName(p1))

	changeNameP(p2)
}

func changeName(p person) person {
	p.first = "Hector"
	return p
}

func changeNameP(p *person) {
	p.first = "Pointer"
	fmt.Println(p)
}
