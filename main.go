package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var joy person
	joy.firstName = "Joy"
	joy.lastName = "Kim"

	fmt.Printf("%+v", joy)
}
