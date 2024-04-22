package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "Jeff"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	//sort.Strings(xs []string) also works
	sort.StringSlice.Sort(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
