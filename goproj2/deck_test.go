package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// In this test, we will be testing the newDeck function:
	// We assign the variable of 'd' to call the function: newDeck()
	// Then we use an if statement, if the length(len) is not equal
	// to 16 (what it should be), then we will return an error (t.Errorf)
	// Use the string and the '%v' value to inject our current length
	// len(d) into the string
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	// This is another if statement to look into the newDeck() function we
	// called and make sure that the first element (card) is an Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	// The d[len(d) - 1] which would pull the last element within the slice
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestADeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected to recieve 16 characters but got %v", d[0])
	}
	if d(len)
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// For the test, we will delete	any files in current directory with the
	// name "_decktesting", then we will create a deck 'newDeck()'
	// Save it to the file, load it from the file and assert the deck length
	// Finally we will delete any files in current directory named "_decktesting"

	// To do this, we will be using the Remove function from the os standard lib
	// Docs: https://pkg.go.dev/os@go1.18.4#Remove
	// Function structure: func Remove(name string) error

	// using 'os.Remove("_decktesting") will be the test file we use
	os.Remove("_decktesting")
	// We are creating a new deck
	deck := newDeck()
	// We are saving this new file to the harddrive
	deck.saveToFile("_decktesting")
	// Now we will attempt to load this from disk from the same file name
	loadedDeck := newDeckFromFile("_decktesting")
	// Now we will be assurting that this has the correct length
	// If the length of loadedDeck isn't equal to 16, throw an error
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	// The last step is to make sure that we have clean the file after we have test
	os.Remove("_decktesting")
}

func TestDealFunc(t *testing.T) {
	cards := newDeck()
	value, remainingCards := deal(cards, 16)
	if len(cards) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(cards))
	}
	value.print()
	remainingCards.print()
}
