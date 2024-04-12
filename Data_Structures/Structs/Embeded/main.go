package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//embedded inner struct => promoted to outer. can access values by agent.first/last/age. Don't have to go agent.person.first
//you cant do agent.person.first if agent already has an agent.first. If you have a field in the outer struct that is the same key as the inner, agent.key will send the one specific to agent. Works with methods as well.
type agent struct {
	person
	ltk      bool
	callSign string
}

func main() {
	a1 := agent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   22,
		},
		ltk:      true,
		callSign: "007",
	}
	a2 := agent{
		person: person{
			first: "miss",
			last:  "Moneypenny",
			age:   29,
		},
		ltk:      true,
		callSign: "008",
	}

	fieldAgents := []agent{a1, a2}

	for _, v := range fieldAgents {
		fmt.Printf("%v\t \n", v.first)
	}

	fmt.Println(a1.person.last, a1.ltk)

}
