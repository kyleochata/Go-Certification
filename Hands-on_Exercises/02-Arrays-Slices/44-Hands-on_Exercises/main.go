package main

import "fmt"

/*
create the following slices from a:
	● [42 43 44 45 46]
	● [47 48 49 50 51]
	● [44 45 46 47 48]
	● [43 44 45 46 47]
*/
func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	b := a[:5]
	c := a[5:]
	d := a[2:7] //last number is exclusive [first index, second index)
	e := a[1:6]
	xxi := [][]int{b, c, d, e}

	fmt.Println(a)
	for _, v := range xxi {
		fmt.Println(v)
	}
}
