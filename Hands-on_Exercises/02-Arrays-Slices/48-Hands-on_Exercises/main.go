package main

import "fmt"

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
*/
func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Money", "I'm 008"}

	xxp := [][]string{jb, mp}

	for v := range xxp {
		fmt.Println(xxp[v])
	}
}
