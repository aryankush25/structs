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
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red": "#ff0000",
	}

	colors["white"] = "#ffffff"
	// delete(colors, "red")

	fmt.Println(colors)

	// p1 := person{firstName: "Alex", lastName: "Anderson", contact: contact{email: "alex@gmail.com", zipCode: 12345}}

	// pointerToPerson := &p1
	// pointerToPerson.updateName("Alexis")

	// p1.updateName("Alexis")

	// fmt.Println(p1)

	// p1.print()
}

func (c contact) print() {
	fmt.Println(c.email, c.zipCode)
}

func (p person) print() {
	fmt.Println(p.firstName, p.lastName)
	p.contact.print()
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}
