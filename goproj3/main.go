package main

import "fmt"

// This creates a type struct that will be used as a type later..
// You have the ability to embbed one struct within another,
// specifically with the person type that is holding the
// value contact with a type of contactInfo
type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

// func main() {
// Here is are delcaring (var) alex of type person:

// 	var alex person

// This is grabbing the value of alex which is of type person and
// appending the first name to the value of Alex and the last name of Anderson:

// 	alex.firstName = "Alex"
// 	alex.lastName = "Anderson"

// This is the print statement to log out the struct and list out
// each field name and the value that is assigned to each one:

// 	fmt.Println(alex)
// 	fmt.Printf("%+v", alex)
// }

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// We are going to place a new struct of contact info below
		// Where we filled the necessary data in the struct to be used
		// You can also call the data type in the struct by using the
		// same struct name if it doesn't have a unqiue name:
		// EX: contact: contactInfo -> contactInfo: contactInfo
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 90201,
		},
	}
	// We can call this function with the value of jim because the recievers put into
	// func print is (p person) in which 'jim' would be a 'person'
	// Below is a good reson for pointers in go, even with the func and
	// calling it to the 'jim' person, it never updated his name to jimmy

	// This is a use case and exmaple for using a Pointer
	// The & is an operating, we place the & and some variable name
	// When we do that, we are saying look at this variable and give us access
	// to the memory address that this variable is pointing at

	// This pointer is not a reference directly to struct, its a reference to the
	// memory address that the struct exists at

	// EDIT II:
	// There also is a trick for writing pointers in go:
	// EX:
	jim.updateName("Jimmy")
	jim.print()
	// In go if we define a receivers with a type of pointer to person
	// When we attempt to call the 'updateName' function go will allow us to
	// call the function with a pointer or the actual type of person

	// The difference is when using jimPointer, this is a pointer to type 'person'
	// When calling 'jim' we are able to get direct access to the TYPE of 'person'
	// Whenever we are using it with a function that has a receiver of the pointer
	// of the same type we are trying to call EX: 'person', we can call it directly
	// as go will turn your varaible of type person into pointer person for you

	// Above is an equal way to use a pointer as the traditional way on the bottom
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()
}

// Whenever we use this star, we place a memory address or the pointer
// Whenever we do that, we are saying take this memory address and give me
// the value that exists within that memory address that holds type person

// Whenever you see a star where the type should be, that means this is a
// description of a type of pointer to "person"
// This update name function can only be called with a reciever of a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// '*pointerToPerson' is the memory address that jim lives at

	// we are saying heres the pointer, I dont want to look at the memory
	// address anymore, give me direct address to the value within it

	// This is an actual operator that takes this pointer '*person' and turns it into a value
	(*pointerToPerson).firstName = newFirstName
}

// This will be a new function that takes structs as recievers
func (p person) print() {
	fmt.Printf("%+v", p)
}

// THE GOLDEN RULE TO POINTERS
// 1.
// If you have a address, you can turn that address into a value with:
// Turn 'address' into 'value' with '*address'
// If you have a value, you can turn that into a address by writing:
// Turn value into address with &value
// 2.
// Understand when we use the *star in front of a type or a memory address
// When you see a star where we would normally specifiy the type, that means:
// We are looking for a type that is a pointer to a person

// SOME GOTCHAS AROUND GO CODE
// For example below is valid code, although go works differently with structs from slices
func GOTCHAS() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

// Exp.
// Here we created a brand new slice and passed it off to the updatedSlice function
// Then we updated the slice and printed out the originial one
// When we printed out this originial slice, we saw that the first element
// Was changed with the updatedSlice function

// This is completely opposite from what we've been seeing with a struct
// Remember, with the struct, when we pass it off to the function,
// go made a copy (interally) of the entire struct
// So when we modified it in the function, we were modifiying the copy, not the orignial struct

// Difference between slices and arrays:

// Arrays:
// Primitive data structure
// Can't be rezied
// Rarely used directly

// Slices:
// Can grow and shrink
// Used 99% of the time for lists of elements

// When creating a slice, go internally is creating two seperate data structures
// when making 'mySlice' go made the slice and array data structures:

// Slice data structure:
// 'pointer to head'
// capacity
// length

// Array data structure:
// It holds that value within the slice: []string{"Hi", "There"}

// When creating a slice in go, go will automatically create the slice data struct
// in one address in memory

// Then that data structure is going to have a pointer to the underlaying array in which
// will be existing at a different memory address
// The 'mySlice' is not pointing at the underlying array
// When referring to mySlice it is actually returning the slice data structure
// not the array
