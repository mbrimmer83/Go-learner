package main

import "fmt"

type contactInfo struct {
	email       string
	phoneNumber int
	zipCode     int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// First way to declare a struct
	// me := person{"Matthew", "Brimmer"}

	// Second way to declare a struct
	// me := person{firstName: "Matthew", lastName: "Brimmer"}

	// Third way to declare a struct
	// var me person
	// me.firstName = "Matthew"
	// me.lastName = "Brimmer"

	// Print formatted struct
	// fmt.Printf("%+v", me)

	// fmt.Printf("My name is %v %v", me.firstName, me.lastName)

	me := person{
		firstName: "Matthew",
		lastName:  "Brimmer",
		contact: contactInfo{
			email:       "mbrimmer1@gmail.com",
			phoneNumber: 7709002043,
			zipCode:     30132,
		},
	}

	// Create pointer to me in order to update firstName
	// mePointer := &me
	// mePointer.updateName("Thomas")

	// Shortcut to pointer from updateName func
	me.updateName("Thomas")
	me.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
