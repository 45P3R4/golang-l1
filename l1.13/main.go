package main

import "fmt"

func main() {
	ac, bc := 1, 2

	//Best solution
	ac, bc = bc, ac
	fmt.Println(ac, bc)

	//XOR (only for integer)
	ax, bx := 1, 2
	ax = ax ^ bx
	bx = ax ^ bx
	ax = ax ^ bx
	fmt.Println(ax, bx)

	//Sum/sub
	as, bs := 1, 2
	as = as + bs
	bs = as - bs
	as = as - bs
	fmt.Println(as, bs)

	//Multiply/divide (not for b=0)
	am, bm := 1, 2
	am = am * bm
	bm = am / bm
	am = am / bm
	fmt.Println(am, bm)
}
