package main

import (
	"encoding/json"
	"fmt"
)

// as long as the fields are capitol, Marshal will work. Ideally the type identifier will be capitol.
type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"Jen","Last":"Nike","Age":33}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person2

	//Unmasrshal takes in a []byte and stores the result in the value pointed to by v json.Unmarshal([]byte, v *any) error
	//err := json.Unmarshal(bs, people) //json: Unmarshal(non-pointer []main.person2)
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
