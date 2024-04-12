/*
add code to print a single value stored in map
use similar code to use lookup of "Q" and print
use "comma ok" to test whether "Q" is stored in the map
	print statement if it is not stored

*/

package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	//print single value stored in the map
	age := m["James"]
	fmt.Println(age)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no key of Q; if key is not found then the zero value of that data-type is returned", v)
	}

	// for k, v := range m {
	// 	fmt.Printf("key: %v \t value: %v\n", k, v)
	// }
}
