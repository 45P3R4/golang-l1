package main

import "fmt"

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice := deleteElement(slice, 5)

	fmt.Println(newSlice)
}
