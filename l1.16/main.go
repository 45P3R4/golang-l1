package main

import "fmt"

func quickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	pivot := numbers[0]

	less := make([]int, 0)
	greater := make([]int, 0)

	for _, num := range numbers[1:] {
		if pivot > num {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = append(quickSort(less), pivot)
	greater = quickSort(greater)
	return append(less, greater...)
}

func main() {
	numbers := []int{12, 20, 22, 60, 1, 5, 10, 80, 18}
	sorted := quickSort(numbers)

	fmt.Println(sorted)
}
