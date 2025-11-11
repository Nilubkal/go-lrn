package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// function extending type deck with a receiver
// (d deck) - receiver !
// Any variable of type 'deck' gets access to the print method.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// _ - dummy iterator with no meaning outside of the iteration.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Defining a function with multiple return values ( GO supports multiple return values from a function)
// returning two values of type deck -> (deck, deck)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// converting a slice of strings into one big string by , separator
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

// Save to local file function following the documentation it needs to return type of error, if something goes wrong.
// the last argument is the permissions of the file if it doesn't exist ( 0666).
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read a file from a the local fs - after you read a file you need to return a string of characters
// No receiver needed here because we need a brand new deck. The func will return a byte slice (bs ) and err (if any or nil)
// Opposite to the WriteFile here we need to go in reverse order from []byte -> string -> []string -> deck
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	// error handling in go like this:
	if err != nil {
		// Option #1 - log the error and return a call to newDeck().
		// Option #2 - stop the execution and log the error.
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
