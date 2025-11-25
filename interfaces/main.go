package main

import "fmt"

// In Go, an interface is a contract that defines a set of methods.
// If a type (like a struct) implements all the methods of an interface,
// it is said to satisfy that interface. This allows functions to accept
// different types that share the same behavior. For a type to satisfy the
// 'bot' interface, it must have a 'getGreeting' method that returns a 'string',
// matching the signature defined below.
type bot interface {
	getGreeting() string
}

// Full interface definition
// To match the interface requirements all the argument types need to be the same everywhere
// type <interfaceName> interface {
// <function-name> (list-of-argument-types: string, int) (list-of-return-types: string, int)

// we are going to declare two custom types as struct - englishBot and spanishBot
type englishBot struct{}
type spanishBot struct{}

func main() {

	// This is how we declare a new value of type struct with no fields associated with it
	eb := englishBot{}
	sb := spanishBot{}

	// without an interface ->
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// we are going to create a printGreeting function that will take an englishBot or a spanishBot and call the getGreeting method on it
// In go if the receiver is not used within the function it can be removed and only the type kept - in both cases in the functions underneath
// the receivers are not used - therefore they can be omitted:
// func (eb englishBot) getGreeting() string { ->
func (englishBot) getGreeting() string {
	//pretending for a very custom logic for english bot
	return "Hi There!"
}

// func (sb spanishBot) getGreeting() string { ->
func (spanishBot) getGreeting() string {
	return "Hola!"
}

// without an interface ->
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) {
// fmt.Println(sb.getGreeting())
// }

// To reduce the usage of two different functions doing the same thing to one we're going to use interfaces
