package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Kyle",
		friends: map[string]int{
			"Dave": 1,
			"Mike": 2,
		},
		favDrinks: []string{"Martini", "Pepsi"},
	}
	fmt.Println(p1)
	fmt.Printf("%v\n%#v\n%T\n", p1, p1, p1)
	for k, v := range p1.friends {
		fmt.Println(p1.first, "-friends-", k, v)
	}
	for _, v := range p1.favDrinks {
		fmt.Println(p1.first, "-drinks-", v)
	}
}
