package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p1 := person{firstName: "Alex", lastName: "Anderson"}

	p1.print()
}

func (p person) print() {
	fmt.Println(p.firstName, p.lastName)
}
