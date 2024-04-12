package main

import (
	"fmt"
	"math/rand"
)

/*
Create a program with sequential and conditional control flow
a) create a random int bw 0 - 250
b) store value of that variable with id of x
c) print out name and value of variable
d) if state ment 0-100 print b/w 0 and 100
*/

// func init is a niladic function ==> Takes no arguements!
func init(){
	fmt.Println("Beep, Boop...Starting up program...PROGRAM ON!")
}

func main() {
	// x := rand.Intn(250)
	// fmt.Printf("x is: %v \n", x)
	// switch {
	// 	case x <= 100:
	// 		fmt.Println("x is between 0 and 100")
	// 	case x >= 101 && x <= 200:
	// 		fmt.Println("x is between 101 and 200")
	// 	case x >= 201 && x <= 250:
	// 		fmt.Println("x is between 201 and 250")
	// 	default:
	// 		fmt.Println("x is not between 0 and 250")
	// }
	for i := 0; i < 11; i++ {
		y := rand.Intn(3)
		fmt.Printf("y is: %v \n", y)
	}
}