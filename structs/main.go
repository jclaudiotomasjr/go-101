package main

import "fmt"

type Pessoa struct {
	Name     string
	LastName string
	Age      uint8
	Status   bool
}

func main() {

	p := Pessoa{
		Name:     "Vitor",
		LastName: "Tomás",
		Age:      13,
		Status:   true,
	}

	p1 := Pessoa{
		Name:     "Yan",
		LastName: "Deslandes",
		Age:      13,
		Status:   true,
	}

	p2 := Pessoa{
		Name:     "Jack",
		LastName: "chefão",
		Age:      4,
		Status:   true,
	}

	fmt.Println(p, p1, p2)

}
