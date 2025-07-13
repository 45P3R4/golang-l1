package main

import "fmt"

func printType(variable any) {
	switch variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan struct{}:
		fmt.Println("chan")
	default:
		fmt.Println("Failed to recognize type")
	}

}

func main() {
	var number int = 1
	var text string = "text"
	var boolean bool = false
	var structure chan struct{} = make(chan struct{}, 0)

	fmt.Println("Recognize type with switch operator:")
	printType(number)
	printType(text)
	printType(boolean)
	printType(structure)

	//If needed only print, not effective on low level
	fmt.Println("\nRecognize type with formatted string:")
	fmt.Printf("%T\n", number)
	fmt.Printf("%T\n", text)
	fmt.Printf("%T\n", boolean)
	fmt.Printf("%T\n", structure)

}
