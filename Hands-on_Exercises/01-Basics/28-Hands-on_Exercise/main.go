package main

import "fmt"

//fxn that takes no arguements ==> niladic
func main() {
	for i := 0; i < 100; i++ {
		// "%" ==> find the remainder (% is also called a modulus)
		if i%2 == 0 {
			//if i / 2 leaves no remainder, the continue reserved word will immediately increase i to the next iteration loop
			continue
		}
		fmt.Printf("%v\n", i)
	}
	fmt.Println("End of all odd numbers")
}
