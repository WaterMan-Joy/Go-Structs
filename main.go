package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	joy := person{
		firstName: "Joy",
		lastName:  "Kim",
		contact: contactInfo{
			email:   "<EMAIL>",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", joy)
}
