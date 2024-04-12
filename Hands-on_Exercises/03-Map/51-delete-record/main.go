package main

import "fmt"

/*
Using the code from the previous example, add a record to your map. Now print the map out
using the “range” loop
key value
`fleming_ian` `steaks`, `cigars`, `espionage
*/
func main() {
	mp := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}
	mp["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	delete(mp, "moneypenny_jenny")
	i := 0
	for k, v := range mp {
		fmt.Printf("%s\t %v\t i: %d\n", k, v, i)
		i++
	}
}
