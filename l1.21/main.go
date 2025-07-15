package main

import "fmt"

type Maid struct{}

func (maid *Maid) Clean() {
	fmt.Println("Cleaning...")
}

type MaidAdapter struct {
	*Maid
}

func (maidAdapter *MaidAdapter) Work() {
	maidAdapter.Clean()
}

func NewMaidAdapter(maid *Maid) Worker {
	return &MaidAdapter{maid}
}

type Worker interface {
	Work()
}

func main() {
	worker := NewMaidAdapter(&Maid{})

	worker.Work()
}
