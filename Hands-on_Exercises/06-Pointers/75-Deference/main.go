package main

import "fmt"

func main() {
	v := "string"
	y := &v

	fmt.Println(y)
	fmt.Println(*y)

}
