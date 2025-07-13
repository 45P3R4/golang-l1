package main

import (
	"fmt"
)

func BinarySearch(slice []int, value int) int {
	low := 0
	high := len(slice)

	for low <= high {
		mid := (low + high) / 2
		guess := slice[mid]

		switch {
		case guess == value:
			return mid
		case guess > value:
			high = mid
		case guess < value:
			low = mid
		default:
			return -1
		}
	}
	return -1
}

func main() {

	const size = 10000000
	var numbers []int = make([]int, 0)
	//Fill slice
	for i := 0; i < size; i++ {
		numbers = append(numbers, i)
	}

	resultIndex := BinarySearch(numbers, 9893958)

	if resultIndex < 0 {
		fmt.Printf("\nIndex not found\n")
	} else {
		fmt.Println("Index:", resultIndex)
	}

}
