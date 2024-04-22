package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"Jeff", 33}
	p2 := person{"Mike", 42}
	p3 := person{"Jen", 25}

	people := []person{p1, p2, p3}
	fmt.Println(people)
	//ByAge(people) == T(n any) ==> TYPE CONVERSION TO BYAGE type (which is a slice of person struct(s))
	fmt.Println("=======byAge==============")
	//Sort(x Interface). type Interface interface{
	// Len() int, Less(i,j) bool, Swap(i,j int)
	//	}
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println(ByAge(people).Len())
	ByAge(people).Swap(0, 1)
	fmt.Println(people)
	fmt.Println("=======ByName==============")
	sort.Sort(ByName(people))
	fmt.Println(people)
}

// ByAge is a slice of person structs. people gets to assume this type because it is a []person
type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

type ByName []person

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByName) Less(i, j int) bool {
	return a[i].age < a[j].age
}
