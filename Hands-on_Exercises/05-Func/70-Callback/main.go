package main

import "fmt"

func main() {
	printSquare(square)
	fmt.Println(printMSquare(mathSquare, 10))
}

func square() string {
	return "i am a square"
}
func printSquare(f func() string) {
	fmt.Println(f())
}
func mathSquare(n int) int {
	return n * n
}
func printMSquare(f func(int) int, num int) string {
	x := f(num)
	return fmt.Sprintf("%v squared is %v", num, x)
}
