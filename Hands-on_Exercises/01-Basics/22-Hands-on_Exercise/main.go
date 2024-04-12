package main

import (
	"fmt"
	"math/rand"
)

/*
Create 2 random ints between 0 inclusive and 10 exclusive
○ assign them to variables with the identifiers x and y
● Print their values
● use an if statement to print the correct description
○ x and y are both less than 4
○ x and y are both greater than 6
○ x is greater than or equal to 4 and less than or equal to 6
○ y is not 5
○ none of the previous cases were met
*/

func main(){
	x := rand.Intn(10)
	y := rand.Intn(10)

	if x < 4 && y < 4 {
		fmt.Printf("x is: %v \n y is: %v \n Both are less than 4", x, y)
		return
	}
	if x > 6 && y > 6 {
		fmt.Printf("x is: %v \n y is: %v \n Both are greater than 6", x, y)
		return
	}
	if x >= 4 && x <= 6 {
		fmt.Printf("x is: %v \n X is greater than or equal to 4, while also being greater than or equal to 6", x)
		if y != 5 {
		fmt.Printf("y is; %v \n Y is not equal to 5", y)
	}
	} else {
		fmt.Printf("x = %v \n y = %v \n No cases matched \n", x, y)
	}
}