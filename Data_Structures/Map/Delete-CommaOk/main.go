package main

import "fmt"

func main() {
	fmt.Println("==============")
	fmt.Println("Make a map")

	mp := make(map[string]int)

	//add el to map: nameOfMap["key"] = value
	mp["todd"] = 42
	mp["mary"] = 35
	mp["greg"] = 80

	fmt.Println(mp)
	fmt.Println("==============")

	//delete(mapName, keyName)
	delete(mp, "mary")

	fmt.Println(mp)

	fmt.Println("===accessing keys that don't exist")
	fmt.Println(mp["mary"]) //won't panic --> gives the zero value of the key's type (int = 0; string = ""; bool = false?)
	/*
		Issue is that we don't actually know if this key exists or it's actual value is the zero value of that key's Type. Use comma ok idiom to handle this
	*/

	// v, ok := mp["mary"]
	// if !ok {
	// 	fmt.Println("no person in map with key of mary")
	// } else {
	// 	fmt.Printf("mary is %d years old", v)
	// }

	//typical way of using the ok. If statement, statement idiom
	//statement1 = v, ok := mp["mary"]
	//statement2 = !ok
	//Reason: it limits the scope of v to just this block. Less chances of accidental reuse
	if v, ok := mp["mary"]; !ok {
		fmt.Println("no person in map with key of mary")
	} else {
		fmt.Printf("mary is %d years old", v)
	}

	fmt.Println("===accessing keys that don't exist")

}
