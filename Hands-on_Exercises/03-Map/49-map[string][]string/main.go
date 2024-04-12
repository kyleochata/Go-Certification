package main

import "fmt"

/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.
key value
`bond_james` `shaken, not stirred`, `martinis`, `fast cars`
`moneypenny_jenny` `intelligence`, `literature`, `computer science`
`no_dr` `cats`, `ice cream`, `sunsets`
*/

func main() {
	mp := map[string][]string{
		"bond_james":       []string{`shaken, not stirred`, `martinis`, `fast cars`},
		"moneypenny_jenny": []string{`intelligence`, `literature`, `computer science`},
		"no_dr":            []string{`cats`, `ice cream`, `sunsets`},
	}
	i := 0
	for k, v := range mp {
		fmt.Printf("%s\t %v\t i: %d\n", k, v, i)
		i++
	}
}
