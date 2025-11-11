// standard Go package library - https://pkg.go.dev/std

// for writing to a local file - we can use os.WriteFile
// func WriteFile(name string, data []byte, perm FileMode) error
// It requires a []byte - byte slice - which means that the input data needs to be converted to slice of bytes !
// asciitable.com - to convert strings into bytes
// We can use type conversion in go -> []byte("Hi There!") <- converting string to bytes

package main

func main() {

	// var card string = "Ace of Spades"
	// the notation underneath is equivalent to the one above !
	// card := "Ace of Spades" // := - this operator will do the typing association and its used the first time its initializing a variable with a type  !
	// if we are re-assigning a new value to an existing variable the =: -> =
	// card = "Five of Diamonds"
	// Another way of assigning the value of a var is by calling a function with a pre-defined type on it
	// card := newCard()

	//cards := deck{"Ace of Diamonds", newCard()}
	// adding an element to the slice
	// append doesn't modify the cards slice but returns a new one by adding the new element !
	// cards = append(cards, "Six of Spades")

	// iteration over a slice of cards
	// range is used to iterate over every single record inside of the slice ( cards )
	// i,card = the index and the actual value being iterated over
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// fmt.Println(cards)

	// calling from the deck.go
	// cards.print()

	// explicit declaration of return type for functions !
	// func newCard() string {
	// 	return "Five of Diamonds"
	// }

	// Working with data structures:
	// Arrays ( Every element in a slice must be of same type !!!)
	// Array - Fixed length list of things
	// Slice - An array that can grow or shrink !
	//		[]string{}  - this is how a slice of strings is being defined,
	//					  all elements of it are inside {}.
	//		slice[startIndexIncluding:upToNotIncluding] - selection in a slice !
	//		slice[:2] - will give 0,1
	//		slice[2:] - will give everything from 2 -> end

	// type - the types are similar to the classes in other programming languages
	// A function with a receiver is like a method - a function that belongs to an instance
	// cards := newDeck()
	// cards.saveToFile("_my_cards.txt")
	//fmt.Println(cards.toString())

	// The deal function can be called without explicit imports because its part of the same package - main
	// hand, remainingCards := deal(cards, 5)

	// we can call the print() function on both of the following vars because they are of type deck
	// hand.print()
	// remainingCards.print()

	// get a deck of cards from a file
	cards := newDeckFromFile("_my_cards.txt")
	cards.print()
}
