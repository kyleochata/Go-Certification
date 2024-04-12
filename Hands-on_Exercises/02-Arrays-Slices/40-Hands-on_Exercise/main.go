package main

import "fmt"

/*
use array literal to store the following elements in the array
let compiler determine len
print out arr and len of arr
*/

//array literal == arr := [n]TYPE{a, b, ...}
//n = number of items in arr; TYPE (data type of n [only 1 type per arr]); {...} = setting values of items in the array according to index of when item is placed in curly

//[...] ==> let the compiler find out the len
func main() {
	flavor := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println("Available flavors are:")
	for v := range flavor {
		fmt.Println(flavor[v])
	}
	fmt.Println("=========================")
	fmt.Printf("There are %v total flavors\n", len(flavor))
	fmt.Println("=========================")
	fmt.Println(flavor)
	fmt.Printf("%T", flavor)
}
