package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Andersen"} <-- works but risky
	// alex := person{firstName: "Alex", lastName: "Andersen"} <-better
	var alex person // populated with 'zero values' aka empty string

	alex.firstName = "Alex"
	alex.lastName = "Andersen"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "Jim@party.com",
			zipCode: 90210,
		},
	}

	// jimPointer := &jim   --- this is longhand version
	// jimPointer.updateName("Jimmmy")
	// short version, GO will auto cast to a pointer for the below func
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
