package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contact
}

func main() {
	p1 := person{firstName: "Alex", lastName: "Anderson", contact: contact{email: "alex@gmail.com", zipCode: 12345}}

	fmt.Println(p1)

	p1.print()
}

func (p person) print() {
	fmt.Println(p.firstName, p.lastName)
	p.contact.print()
}

func (c contact) print() {
	fmt.Println(c.email, c.zipCode)
}
