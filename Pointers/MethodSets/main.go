package main

import "fmt"

type dog struct {
	first string
}

//Methods are available based on type. If type T then all methods with receiver typeT; if type *T then all methods for type T and *T
func (d dog) walk() {
	fmt.Println("My name is ", d.first, "and I'm walking")
}
func (d *dog) run() {
	d.first = "Speed"
	fmt.Println("My name is", d.first, "Zoooooommm!")
}

//because run has a pointer receiver => goodDog must be a value of pointer to implement the goodDog interface.
//When implementing interfaces if it's a type *T then you can use methods with receivers for type *T and T for definition. if method set is of type T then you can only define the interface with type T methods.
type goodDog interface {
	walk()
	run()
}

func goodRun(g goodDog) {
	g.run()
}

func main() {
	d1 := dog{"Mary"}
	d2 := &dog{"Hiccup"}
	//Calling functions is fine. It matter when it comes to invoking interface functions/methods
	d1.run()
	d1.walk()
	//goodRun(d1) //cannot use d1 (variable of type dog) as goodDog value in argument to goodRun: dog does not implement goodDog (method run has pointer receiver)

	goodRun(d2) //works because it's a pointer
	d2.run()
	d2.walk()
}
