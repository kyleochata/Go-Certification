package main

import (
	"fmt"
	"sync"
)

// Method set => Wait (wg *sync.WaitGroup)
// Wait() has a receiver that's a pointer. However we created a non-pointer variable in wg => using the Wait() method eventhough wg is not a pointer
// if a (t T) receiver uses T and *T values. regular and pointer types to invoke the method. (t *T) receiver uses only *T. Pointer types are only ones allowed to invoke a method with the (t *T) receiver
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go foo()
	bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

/*
Why This Code Works
In your example, wg is a regular sync.WaitGroup instance, not a pointer. However, due to Go's method set rules, the Add, Done, and Wait methods can still be called on it.
Go provides automatic conversion from a regular type to its pointer counterpart when necessary. This feature, known as method promotion, allows a regular type (sync.WaitGroup) to call methods that might be defined with pointer receivers (*sync.WaitGroup). This is why wg.Add(2), wg.Done(), and wg.Wait() work, even if wg isn't explicitly defined as a pointer.
Summary
Your code works because Go allows a defined type (like sync.WaitGroup) to call methods with either a regular or pointer receiver, thanks to its method set rules and method promotion capabilities. This seamless handling of regular and pointer receivers is one of Go's design features that simplifies code usage and reduces the need for explicit type conversion.

// Define a custom struct
type Person struct {
	Name string
	Age  int
}

// Method with a regular receiver
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// Method with a pointer receiver
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Println(p.Name, "just had a birthday! Now they're", p.Age)
}

func main() {

	// Create an instance of Person (not a pointer)
	p := Person{Name: "Alice", Age: 25}

	// Call the method with a regular receiver
	p.Greet() // This works because Greet has a non-pointer receiver

	// Call the method with a pointer receiver
	// Go automatically converts 'p' to '*Person' to call HaveBirthday
	p.HaveBirthday() // This works because of method promotion

	// Create a pointer to Person
	p2 := &Person{Name: "Bob", Age: 30}

	// Call the method with a pointer receiver
	p2.HaveBirthday() // This works because HaveBirthday has a pointer receiver

	// To call the method with a regular receiver, you need to dereference the pointer
	(*p2).Greet() // This works because we've dereferenced p2 to get a Person
}


*/
