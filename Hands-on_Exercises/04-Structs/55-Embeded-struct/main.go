package main

import "fmt"

/*
● Create a type engine struct, and include this field
	○ electric bool
● Create a type vehicle struct, and include these fields
	■ engine
	■ make
	■ model
	■ doors
	■ color
● Create two VALUES of TYPE vehicle
	○ use a composite literal
● Print out each of these values.
● Print out a single field from each of these values.
*/

type engine struct {
	electric bool
}
type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	aR8 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Audi",
		model: "R8",
		doors: 2,
		color: "Silver",
	}

	a4 := vehicle{engine{true}, "Audi", "R8", 2, "Silver"}

	aCarList := []vehicle{a4, aR8}

	for k, v := range aCarList {
		fmt.Printf("%v\n %T\t %#v\n", k, v, v)
	}
}
