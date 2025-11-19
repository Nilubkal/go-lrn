// Struct - Data structure - Collection of properties that are related together.( could be similar to a Dict in Python)
package main

import "fmt"

// when we create a struct first we define all the properties that it will have and then we create a value that matches this definition
// for the tutorial we are going to create a person struct. First we are going to define all the fields that define a person (with type of these fields)
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo // it can be just contactInfo without adding a new field contact
	contactInfo
}

func main() {

	// Making a new person - direct way of declaring a struct
	// one other way is alex := person{"Alex", "Anderson"} - this way you relay on the order of definition of fields - not good !
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// Another way of declaring a struct
	//var alex person // this declares a value of type person - go assigns a 0 value before populating it
	// Default values based on the type:
	// string -> ""
	// int -> 0
	// float -> 0
	// bool -> false
	// Declaring alex firstName and lastName
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//fmt.Println(alex)
	// %+v prints out the field names and values of the struct
	//fmt.Printf("%+v", alex)

	// Embedding one struct inside of another
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// this is defining explicitly a new field contact
		// contact: contactInfo{
		// whenever you are defining a multiline struct like this, every single line needs have a ,
		// even the last one after the closing }!
		// email:   "jim@gmail.com",
		// zipCode: 94000,
		// it can be done like this also by skipping the new field definition from the struct
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// &var -> operator - Give me the memory address of the value this variable is pointing at
	// jimPointer := &jim // memory address == pointer
	// jimPointer.updateName("Jimmy")

	// less code doing the same as using the jimPointer hop from above
	// jim.updateName("Jimmy"): Even though jim is a value,
	// Go automatically takes its address (&jim) and passes it as a pointer to the updateName method
	// because the method is defined with a pointer receiver. This is a convenience feature of Go.
	// In other words if you have a variable of the same type as the pointer that acts as a receiver
	// GO will translate the memory allocation with pre-assignment using the & since they are of the same struct type.
	jim.updateName("Jimmy")
	jim.print()
}

// struct can have receivers too
func (p person) print() {
	fmt.Printf("%+v", p)
}

// Function that will update the first name of the receiver
// Without using the concept of pointers in go - this won't modify the first name and it will return:
// {firstName:Jim lastName:Party contactInfo:{email:jim@gmail.com zipCode:94000}}
//func (p person) updateName(newFirstName string) {
//	p.firstName = newFirstName
//}

// Function that will actually update the first name using pointers
// Since Go is pass by value language.
// *pointer - give me the value this memory address is pointing at
// *person - this is a type description - it means we are working with a pointer to a person
// Additional explanation:
//
//			func(pointerToPerson *person) updateName(){
//				(*pointerToPerson).firstName = newFirstName
//				}
//	     *person -> this is a type description it means we're working with a pointer to a person
//	     *pointerToPerson -> this is an operator -> give me the value this memory address is pointing at.
//
// This means that the function can be called with a receiver which is a pointer to a type person , and then this pointer will be manipulated.
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Rule for using * and &
// Turn memory address into value  - > *address
// Turn value into memory address - > &value

// Slices vs Arrays & Why Slices Don't Need Pointers for Modification

// An Array in Go is a fixed-length sequence of items of the same type.
// The length is part of the array's type. For example, [5]int and [10]int are different types.
// Because of their fixed size, they are not used as often as slices.
// Example: var myArray [5]int

// A Slice, on the other hand, is a dynamically-sized, flexible view into the elements of an underlying array.
// A slice is a data structure (a struct, internally) that holds three pieces of information:
// 1. A pointer to an underlying array where the data is stored.
// 2. The length of the slice (number of elements it contains).
// 3. The capacity of the slice (number of elements in the underlying array from the start of the slice).

// Why can a slice be modified by a function without a pointer?
// When you pass a slice to a function, a copy of the slice header (the pointer, length, and capacity) is made.
// However, the pointer inside this new copied slice header still points to the *same* underlying array.
// Therefore, if the function modifies the elements of the slice (e.g., mySlice[0] = "new value"), it is actually modifying the data in the original underlying array.
// This is why the changes are visible outside the function, even though the slice itself was passed by value. It's a reference type.
// Things that you shouldn't worry using pointers in GO are:
// Reference types:
//   slices
//   maps
//   channels
//   functions
// Things that you should worry using pointers:
// Value Types:
//   structs
//   int
//   float
//   bool
//	 string
