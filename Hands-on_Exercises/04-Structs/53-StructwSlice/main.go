package main

import "fmt"

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
*/

type person struct {
	first       string
	last        string
	favIceCream string
}

func main() {
	me := person{"kyle", "smith", "vanilla"}
	jen := person{"jen", "smith", "vanilla"}

	fmt.Println(me.first, me.last, me.favIceCream)
	fmt.Println(jen.first, jen.last, jen.favIceCream)

}
