package main

import "fmt"

func main() {
	//anon fxn  func(p param(s))(return(s)){<code>}(a arg(s))

	//function expression = functions has the same rights as any other Types in Go. Able to assign functions to variables, passed as args to fxns, used in control flow statements (if, for) ==> first class citizen
	//Expression = combo of values, variables, operators, and fxn calls that're evaluated to produce a single value.
	x := func(s string) {
		fmt.Println("General ", s)
	}
	x("Kenobi")
}
