package main

import (
	"strconv"
)

func setBitToOne(number int64, bit int64) int64 {
	// return number | bit
	return number | (1 << bit)
}

func setBitToZero(number int64, bit int64) int64 {
	return number ^ (1 << bit)
}

func main() {
	var num int64
	num = 8
	println(strconv.FormatInt(int64(num), 2))

	num = setBitToOne(num, 1)
	println(strconv.FormatInt(int64(num), 2))

	num = setBitToZero(num, 1)
	println(strconv.FormatInt(int64(num), 2))

}
