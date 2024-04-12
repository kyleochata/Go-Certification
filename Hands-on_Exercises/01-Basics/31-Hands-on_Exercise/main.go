package main

import "fmt"

/*
use a map provided
for range loop to print value and key of associated value
*/

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("key: %v \t value: %v\n", k, v)
	}
}
