package main

import "fmt"

func main() {
	a := 171
	evv, odd := evenOrOdd(a)
	fmt.Println("Evens", evv)
	fmt.Println("Odds", odd)
	fmt.Println("Num of Evens", len(evv))
	fmt.Println("Num of Odds", len(odd))
}
