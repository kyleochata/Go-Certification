package main

import "fmt"

/*
make a loop that runs 5x
nest a loop within that runs 5x
print something in each iteration of ea loop
*/
func main() {
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Printf("outer loop iteration: %v\t", i)
			fmt.Printf("inner loop iteration: %v\t", j)
			if j%2 == 0 {
				fmt.Printf("this iteration is an even number")
			}
			fmt.Println()
		}
	}
}
