// We will always start out with package main
package main

// This imports any package wether it's from the standard library or not
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}
	// For loops with the '_' is a blank arguement that doesn't need to be called
	for _, suit := range cardSuit {
		for _, value := range cardValue {
			// The append built-in function appends elements to the end of a slice.
			// If it has sufficient capacity, the destination is resliced to accommodate
			// the new elements. If it does not, a new underlying array will be allocated.
			// Append returns the updated slice. It is therefore necessary to store the
			// result of append, often in the variable holding the slice itself:
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// The [:] is used for how you want it to be called
// Ex: [:Dominic] === < [Dominic:] === >
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Using import package called 'strings' to join a slice of strings
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Front parenthesis is reciever, on function is arguements, end is returned
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 = log the error and entirely quit the program
		// We will be going with option 2 while using the os package

		// OS Package: Package os provides a platform-independent
		// interface to operating system functionality. The design is
		// Unix-like, although the error handling is Go-like; failing
		// calls return values of type error rather than error numbers.
		// Often, more information is available within the error.
		// For example, if a call that takes a file name fails,
		// such as Open or Stat, the error will include the failing
		// file name when printed and will be of type *PathError,
		// which may be unpacked for more information.

		fmt.Println("Error: Something has gone wrong", err)
		os.Exit(1)
	}
	// Using the function split from the "strings" package in
	// Standard Library,

	// EX: func Split(s, sep string) []string

	// Split slices s into all substrings separated by sep and
	// returns a slice of the substrings between those separators.
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		// The newPosition is going to be the Intn function inside of rand library
		// We want to generate a number between 0 and the legnth of the slice - 1
		// Intn returns, as an int, a non-negative pseudo-random number in the half-open
		// interval [0,n). It panics if n <= 0.
		newPosition := r.Intn(len(d) - 1)
		// This is our exchange logic, which is how we are going to shuffle the cards
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
