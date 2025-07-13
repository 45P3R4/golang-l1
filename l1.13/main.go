package main

import "fmt"

func main() {
	firstNum := 1
	secondNum := 2

	fmt.Println(firstNum, secondNum)

	firstNum, secondNum = secondNum, firstNum

	fmt.Println(firstNum, secondNum)
}
