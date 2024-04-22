package main

import (
	"encoding/json"
	"fmt"
)

// If the fields aren't capitol, the Marshal won't be able to work. Will return empty [{}, {}].
type person struct {
	first string
	last  string
	age   int
}

// as long as the fields are capitol, Marshal will work. Ideally the type identifier will be capitol.
type person2 struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		first: "James",
		last:  "Smith",
		age:   40,
	}
	p2 := person{
		first: "Mike",
		last:  "Jones",
		age:   34,
	}
	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Println("==============")

	p3 := person2{
		First: "Jen",
		Last:  "Nike",
		Age:   33,
	}

	group2 := []person2{p3}
	bs2, err2 := json.Marshal(group2)
	if err2 != nil {
		fmt.Println(err2)
	}

	//have to convert back to string to display. marshal returns []byte
	fmt.Println(string(bs2))
}
