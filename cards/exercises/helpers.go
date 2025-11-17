package main

// func to iterate trough a slice of numbers between 0 and 10
// and print even and odd ones

func evenOrOdd(n int) (_even []int, _odd []int) {

	for d := range n {
		if d != 0 {
			if d%2 != 0 {
				_odd = append(_odd, d)
			} else {

				_even = append(_even, d)
			}
		}
	}
	return
}
