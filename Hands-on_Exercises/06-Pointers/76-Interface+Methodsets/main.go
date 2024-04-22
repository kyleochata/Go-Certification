package main

import "fmt"

type engine struct {
	cylinders int
	size      string
}

func (e *engine) vroom() {
	fmt.Println("Vrooom vroom!")
}

type gasGuzzler interface {
	vroom()
}

func driveTheCar(g gasGuzzler) {
	g.vroom()
}

func main() {
	e1 := &engine{
		cylinders: 8,
		size:      "3.2L",
	}
	e2 := engine{
		cylinders: 6,
		size:      "2.0L",
	}

	driveTheCar(e1)
	driveTheCar(&e2)
}
