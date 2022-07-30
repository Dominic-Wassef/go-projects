package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	email     string
	description
}

type description struct {
	age     int
	zipCode int
	city    string
}

func main() {
	person1 := person{
		firstName: "Jim",
		lastName:  "Morris",
		email:     "jimmorris@gmail.com",
		description: description{
			age:     32,
			zipCode: 90210,
			city:    "LA",
		},
	}
	person2 := person{
		firstName: "Marcus",
		lastName:  "Johnson",
		email:     "MarcusJ@gmail.com",
		description: description{
			age:     63,
			zipCode: 32601,
			city:    "Gainesville",
		},
	}
	person1Pointer := &person1
	person2Pointer := &person2
	person1Pointer.updateAgeZipcodeCity(66, 34639, "Tampa")
	person2Pointer.updateAgeZipcodeCity(77, 00000, "N/A")
	print(person1)
	print(person2)
}

func (pointerToPerson *person) updateAgeZipcodeCity(newAge int, newZipcode int, newCity string) {
	(*pointerToPerson).age = newAge
	(*pointerToPerson).zipCode = newZipcode
	(*pointerToPerson).city = newCity
}

func print(p person) {
	fmt.Printf("%+v", p)
}
