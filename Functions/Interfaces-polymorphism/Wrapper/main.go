package main

/*
type Stringer interface {
	String() string
}
Stringer interface is implemented by any Value that has the method String()
	String() - used to print values passed as an operand to any format that accepts a string or to an unformatted printer such as Print.
*/

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

// attach method to the struct via the receiver --> Stringer interface is active now for book types
func (b book) String() string {
	return fmt.Sprint("this is the book ", b.title)
}

type count int

func (c count) String() string {
	//c is of Type count --> count has an underlying type of int --> strconv.Itoa(i int) --> if you plug c in right away, it will throw a type error because c is currently the type of "count". Have to T() to satisfy the Itoa method (int to asci)
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

// logInfo = wrapper fxn
func logInfo(s fmt.Stringer) {
	log.Println("From logInfo()", s.String())
}

func main() {
	// b & c have their own type (book, count) as well as the Stringer type. Polymorphism!
	b := book{
		title: "Harry Potter",
	}

	var n count = 50

	logInfo(b)
	logInfo(n)

}
