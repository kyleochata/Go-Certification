package main

import "fmt"

func main() {
	var i int

	// for i <= 5 {
	// 	fmt.Printf("first loop: %v \t", i)
	// 	for j := 1; j <= 5; j++ {
	// 		if j == 4 {
	// 			break
	// 		}
	// 		fmt.Printf("inner loop: %v \t \n", j)
	// 	}
	// 	i++
	// }
	for {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}

// the for loop will run until i > 5, once it reaches the break, it escapes the for loop!
