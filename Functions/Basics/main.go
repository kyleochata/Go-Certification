package main

import "fmt"

func main() {
	defer test() //run after main() completes rest of logic
	singleP("Mike")
	s := oneNone("Physics")
	fmt.Println(s)

	s2, n := twoTwo("Kyle", 30)
	fmt.Println(s2, n)
}

/*
func (r receiver) identifier(parameters) (return){ <code> }
	Define our func w/ parameters (if any)
	Call our func and pass in arguements (if any)
Everything in Go is Pass By VALUE
*/
//no params, no returns
func test() {
	fmt.Println("from test func")
}

//1parm, 0 ret
func singleP(s string) {
	fmt.Printf("%v is speaker of today's lecture\n", s)
}

//1parm, 1 ret
func oneNone(s string) string {
	return fmt.Sprintf("%v is the topic for today\n", s)
}
func twoTwo(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years:"), age
}
