//Callback => when you pass a func in as an arg for another func

package main

import "fmt"

func main() {

	fmt.Println(mathIsMathin(4, 3, add))
	fmt.Println(mathIsMathin(4, 3, sub))
	fmt.Printf("%T\n ", mathIsMathin)

}

func mathIsMathin(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
