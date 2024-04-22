package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13, 23, 34, 59, 198, 238, 38, 49, 49, 20, 8}
	xi2 := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13, 23, 34, 59, 198, 238, 38, 49, 49, 20, 8}
	//xi2 := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	fmt.Println("=======sort.Sort()=======")
	start := time.Now()
	sort.Ints(xi)
	fmt.Println(xi)
	duration := time.Since(start).Nanoseconds() / 1000 // Convert to microseconds
	fmt.Printf("Duration: %d microseconds\n", duration)
	fmt.Println("=======bubble sort========")
	start = time.Now()
	bubble_sort(xi2)
	fmt.Println(xi2)
	speed := time.Since(start).Nanoseconds() / 1000 // Convert to microseconds
	fmt.Printf("Speed: %d microseconds\n", speed)

	// fmt.Println(xs)
	// // sort xs
	sort.StringSlice.Sort(xs)
	fmt.Println(xs)
}

func bubble_sort(xi []int) []int {
	for i := 0; i < len(xi); i++ {
		for j := 0; j < len(xi)-1-i; j++ {
			if xi[j] > xi[j+1] {
				xi[j], xi[j+1] = xi[j+1], xi[j]
			}
		}
	}
	return xi
}
