package main

import "fmt"

func main() {
	// map[Type for keys] Type for values
	m := map[string]int{
		"todd": 42,
		"mary": 25,
	}

	fmt.Println(m)
	fmt.Printf("%#v\n", m)
	fmt.Printf("%T\n", m)
	// // ==== OR ====
	fmt.Println("==============")
	fmt.Println("Make a map")
	mp := make(map[string]int)

	//add el to map: nameOfMap["key"] = value
	mp["todd"] = 42
	mp["mary"] = 35

	fmt.Println(m)
	fmt.Println(mp)
	fmt.Println("==============")

	for k := range m {
		fmt.Printf("just the keys: %s\n", k)
	}
	for k, v := range m {
		fmt.Printf("%s is %d years old\n", k, v)
	}
	for _, v := range m {
		fmt.Printf("just the values: %d\n", v)
	}
	if v, ok := m["Padget"]; ok {
		fmt.Printf("%s is %d years old\n", "Padget", v)
	} else {
		fmt.Printf("%s is not found\n", "Padget")
	}

}
