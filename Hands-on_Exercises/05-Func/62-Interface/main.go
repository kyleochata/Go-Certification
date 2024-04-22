package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (s square) Area() float64 {
	return s.length * s.width
}
func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	Area() float64
}

func Info(s shape) {
	fmt.Println(s.Area())
}

func main() {
	s1 := square{
		length: 44,
		width:  10,
	}
	c1 := circle{
		radius: 10,
	}
	Info(s1)
	Info(c1)
}
