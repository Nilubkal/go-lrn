// test go file for the cards project
// Tests for go are usually written separately in go , no additional frameworks are used
// It's accepted for them to end with <porject>_test.go
// Rule of thumb is for each function there should be a test case for it too. Or in general test whatever makes sense.

package main

import (
	"os"
	"testing"
)

// This function will be auto called by the go testing runner with an argument t with type *testing.T
func TestNewDeck(t *testing.T) {
	// Test the length of a Deck
	// t - test handler
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Here we are checking the order of the deck by taking the first and last cards.
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// Tests for the save file and load file functions
// Delete files in the current directory with the name *_decktesting to remove any leftovers
// test the main functionality create - save - load and delete again *_decktesting files.

// Good practice is to name the test functions based on the original functions by adding a prefix of Test<original-func-name>
func TestSaveToDeckAndNewDeckFromFIle(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
