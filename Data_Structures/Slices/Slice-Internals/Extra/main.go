package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}
	g := []float64{3, 1, 4, 2}

	fmt.Println(medianOne(n))
	fmt.Println("after medianOne", n)
	fmt.Println(medianTwo(g))
	fmt.Println("after medianTwo", g)
}

/*
uses the same underlying array --> arr is passed by value (all of go)
the value of n is being passed into the func ==> assigned to variable x
x is pointing to the same underlying array as n
*/
func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(x []float64) float64 {
	// allocate a new underlying array
	n := make([]float64, len(x))
	copy(n, x) //give new n array the values of the x arr

	sort.Float64s(n)
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
	}
	return (n[i-1] + n[i])
}

//medianTwo has a new slice with a different pointer than the slice that was fed into medianTwo. That is why when we look up array "g" after medianTwo is run, that arr "g" will not be sorted and will be in it's original state and order
