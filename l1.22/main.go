package main

import (
	"fmt"
	"math"
	"math/big"
)

func sum(a int, b int) (int64, big.Int) {
	if math.MaxInt-a < b {
		var result big.Int
		result.Add(big.NewInt(int64(a)), big.NewInt(int64(b)))
		return -1, result
	} else {
		result := a + b
		return int64(result), *big.NewInt(-1)
	}
}

func main() {
	a := math.MaxInt
	b := math.MaxInt

	resI, resB := sum(a, b)

	if resI != -1 {
		fmt.Println(resI)
	} else {
		fmt.Println(&resB)
	}
}
