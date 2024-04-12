package main

import "fmt"

/*
Multidimensional slice is like a spreadsheet
	a slice of a slice of some type
*/
func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jp := []string{"Jenny", "Penny", "Rum", "Cherry"}

	// fmt.Println(jb)
	// fmt.Println(jp)

	//slice to store people
	/*
		first [] => signals we would like a slice
		[]string => giving the slice the Type = []string;
		A slice to store slices of strings
	*/
	xxs := [][]string{jb, jp}
	for _, v := range xxs {
		fmt.Println(v)
	}
	fmt.Println(xxs)
}

/*
slices will point to their underlying array
	pointers show where the values are stored and telling the var or func to go to that address for the values
*/
