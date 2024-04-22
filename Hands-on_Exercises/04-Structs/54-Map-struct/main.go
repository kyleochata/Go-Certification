package main

import "fmt"

/*
store the VALUES of type person in a map with
the KEY of last name. Access each value in the map. Print out the values, ranging over the
slice.

*/
// fmt.Println(me.first, me.last, me.favIceCream)
// fmt.Println(jen.first, jen.last, jen.favIceCream)

// people := []person{me, jen}

// for _, v := range people {
// 	fmt.Printf("%v\t %v\n", v.first, v.last)
// }
type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {

	me := person{"kyle", "smith", []string{"vanilla", "birthday cake"}}
	jen := person{"jen", "smith", []string{"vanilla", "birthday cake"}}

	mp := map[string]person{
		me.last:  me,
		jen.last: jen,
		/*me and jen have the same last name of "smith". map is first set as {
			"smith": me{kyle, smith, {vanilla, birthday cake}}
		when jen is added to the map, "smith" is jen last field value. This will overwrite the previous "smith" key that was set to me's value. So when you map, only jen smith values will be in the mp */
		me.first: me,
		//getting the me struct into the map.
	}

	for _, v := range mp {
		fmt.Println(v)
		for _, v2 := range v.favIceCream {
			fmt.Println(v.first, v.last, v2)
		}
	}

}
