package main

import "fmt"

/*
slice l8iteral to store elements in a slice => print out slice and len of slice
*/

func main() {
	tastes := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	fmt.Printf("There are %v total flavors\n", len(tastes))
	fmt.Println("=========================")
	fmt.Println("Available flavors are:")
	// for range will return two values; blank identifier is throwing one away
	for _, v := range tastes {
		fmt.Println(v)
	}
	/*
		for v:=range tastes {
			fmt.Println(tastes[v])
		}
	*/

	fmt.Println("=========================")
	fmt.Println(tastes)
	fmt.Printf("%T", tastes)
}
