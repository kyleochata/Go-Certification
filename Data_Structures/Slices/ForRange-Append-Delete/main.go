package main

import "fmt"

func main() {

	/*
		slices are similar to arrays. They are shaped like them and can only hold values of the same type.
			Unlike arrs, slices will allow you to change the size of the slice where as arrays in Go(lang) are static and unchanging in number of items
	*/

	xi := []int{10, 9, 8, 7}

	//for range loop
	//for index, value := range arr {}
	// for range loops will return an index and value, use blank identifier if first value is not needed
	// for i, v := range xi {
	// 	fmt.Printf("position: %v \t value: %v\n", i, v)
	// }
	// fmt.Println("==============================")
	// for _, v := range xi {
	// 	fmt.Println(v)
	// }

	// fmt.Println("==============================")
	/*
		APPEND to SLICE (adding)
		adding to a slice ==> append(slice []Type, elems ...Type) []Type
		first arg is always the slice you want to append
		Following the first are are any number of elements to be added to the slice as long as they are the correct TYPE for the slice (...Type == variadic parameter; can have unlimited amount)
		returns slice back
		bc append return a slice of the same Type back, we can assign it to the original variable for the first slice
	*/
	fmt.Println("Append to slice")
	xi = append(xi, 10)
	for i, v := range xi {
		fmt.Printf("position: %v \t value: %v\n", i, v)
	}
	fmt.Println("==============================")
	/*
		SLICE a SLICE (cut away)
		Slicing a slice: cutting parts of the slice away = colon operator\
		all of these look like they are based off of array index position (0index)
	*/
	fmt.Println("Slice a slice")
	xi2 := []int{0, 1, 2, 3, 4, 5, 6}

	//[inclusive:exclusive]
	fmt.Println("----------------------")
	fmt.Println("[inclusive:exclusive]")
	fmt.Printf("xi -%#v\n", xi2[0:4])
	fmt.Println("----------------------")
	fmt.Println("[:exclusive]")
	fmt.Printf("xi -%#v\n", xi2[:2])
	fmt.Println("----------------------")
	fmt.Println("[inclusive:]")
	fmt.Printf("xi -%#v\n", xi2[3:])
	//start at position 3 and give all items until the end

	fmt.Println("----------------------")
	fmt.Println("[:]")
	fmt.Printf("xi -%#v\n", xi2[:])
	//give all; from start to end
	fmt.Println("==============================")
	/*
		To delete from a slice, multiple appends must be used!
	*/
	fmt.Println("Delete from slice")
	xiDelete := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi -%#v\n", xiDelete)
	fmt.Println("----------------------")
	//xiDelete[5:] doesn't work due to the append fxn
	// xiDelete[5:] is a slice Type. append takes (slice, elements ...Type) => we are trying to put xiDelete into the second arg where individual elements should be going
	//by adding xiDelete[5:]..., we are spreading out the slice so that it's like we are plugging in the rest of the elements we want (value at index 5 until the end)
	xiDelete = append(xiDelete[:4], xiDelete[5:]...)
	fmt.Printf("xi -%#v\n", xiDelete)

}
