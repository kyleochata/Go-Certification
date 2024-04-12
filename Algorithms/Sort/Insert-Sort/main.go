package main

import "fmt"

func insertionSort(arr []int) {
	for hole := 1; hole < len(arr); hole++ {
		removedVal := arr[hole]
		leftOfRemoved := hole - 1

		//leftOfRemoved >= 0 --> don't go past the beginning of the slice
		for leftOfRemoved >= 0 && arr[leftOfRemoved] > removedVal {
			//give the hole the value of the left position
			arr[leftOfRemoved+1] = arr[leftOfRemoved]
			//lor moves one to L
			leftOfRemoved = leftOfRemoved - 1
		}
		//set hole (should be at the position where all el to L are smaller, and all el to R are larger) to the value of the removed value
		arr[leftOfRemoved+1] = removedVal
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Unsorted array is:", arr)
	insertionSort(arr)
	fmt.Println("Sorted array is:", arr)
}
