package main

import "fmt"

//Main displays numbers 1-9 via a for loop
func main() {
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
}
