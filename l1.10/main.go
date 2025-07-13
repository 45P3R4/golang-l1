package main

import "fmt"

func main() {
	var temrepatures = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempByGroup := make(map[int][]float32)

	for _, temp := range temrepatures {
		group := int(temp) - int(temp)%10

		val, ok := tempByGroup[group]
		if ok {
			tempByGroup[group] = append(val, temp)
		} else {
			tempByGroup[group] = []float32{temp}
		}
	}

	fmt.Println(tempByGroup)
}
