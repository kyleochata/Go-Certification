package main

import "fmt"

func main() {
	fmt.Println(true && true)  // true both sides of && are truthy values
	fmt.Println(true && false) // false: L == truthy, R == falsey
	fmt.Println(true || true)  //true: is L || R truthy? aslong as one is then show true
	fmt.Println(true || false) // true
	fmt.Println(!true)         //give me the opposite of this value (truthy -> falsesy and vice versa)
}
