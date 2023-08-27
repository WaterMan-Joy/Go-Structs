package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	joy := person{
		firstName: "joy",
		lastName:  "Kim",
		contactInfo: contactInfo{
			email:   "<EMAIL>",
			zipCode: 12345,
		},
	}
	// joy.updateNameCopy("Han", "Y")
	// fmt.Println(&joy.firstName)
	// joyPointer := &joy
	joy.updateName("Sunny", "J")
	joy.print()
	fmt.Println(&joy.firstName)
}

func (p person) updateNameCopy(firstName, lastName string) {
	p.firstName = firstName
	p.lastName = lastName
}

func (pointerToPoerson *person) updateName(firstName string, lastName string) {
	(*pointerToPoerson).firstName = firstName
	(*pointerToPoerson).lastName = lastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
