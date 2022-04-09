package main

import "fmt"

type contactInfo struct {
	email string
	phone int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "John",
		lastName:  "doe",
		contactInfo: contactInfo{
			email: "john@example.com",
			phone: 9099081202,
		},
	}

	alex.updateName("mamad")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
