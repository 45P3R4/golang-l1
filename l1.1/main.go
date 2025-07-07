package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Greetings() {
	fmt.Println("Hello! My name is:", h.Name)
}

type Action struct {
	ActionName string
	Human
}

func (a Action) Info() {
	fmt.Println(a.Human.Name, "is", a.ActionName)
}

func main() {
	action := Action{
		ActionName: "Sleeping",
		Human: Human{
			Name: "Maria",
			Age:  26,
		},
	}

	action.Greetings()
	action.Info()
}
