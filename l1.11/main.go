package main

import (
	"fmt"
)

func findIntersection(a []int, b []int) (intersections []int) {
	intersectionMap := make(map[int]any)
	for _, num := range a {
		_, ok := intersectionMap[num]
		if !ok {
			intersectionMap[num] = struct{}{}
		}
	}

	for _, num := range b {
		_, ok := intersectionMap[num]
		if ok {
			intersections = append(intersections, num)
		}
	}

	return intersections

}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	intersections := findIntersection(a, b)

	fmt.Println(intersections)

}
