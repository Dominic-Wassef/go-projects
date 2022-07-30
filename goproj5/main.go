package main

import "fmt"

// This creates a new custom type interface, this is great for reusing code
// If you are a type in this program with a function call 'getGreeting()' and you
// return a string then you are now an honorary member of type 'bot'
// Now that you're also an honorary member of type 'bot', you can now call this
// function called 'printGreeting()'.

// Because of this, go is able to link the bot interface and the
// englishBot / spanishBot structs, again because it matches the coniditions
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	// This is how we declare the value of a empty struct
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// When working with interfaces, use the interface as an arugement to the function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// NOTE: In go, you can create two function with the same
// name so long as they have a different recievers
// Although this rule only applys to recievers, not arugements or return values
func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "Hola!"
}

// Example of more complex interfaces
// Here we are saying that the bot interface should have the getGreeting func
// With two arugements of (string, int) along with two returns of (string, error)
// IMPORTANT NOTE: With interface, you do not need to specify the reciever
// AND... You can have multiple functions within an interface
// type bot interface {
// 	getGreeting(string, int) (string, error)
// }

// There two groups of different types!
// Concrete types: map, struct, int, string and custom types such as 'englishBot'
// Interface types: bot

// Interfaces are 'implicit' meaning we don't have to
// manually say that our custom type satisfies some interface

// Interfaces are a contact to help us manage types: GARBAGE IN -> GARBAGE OUT
// If our custom types implementation of a function is broken, interfaces wont help
