package main

import "fmt"

func main() {
	const (
		FIRST   = iota           // iota = 0, FIRST = iota
		SECOND                   // iota = 1, SECOND = FIRST (=iota, =1)
		THIRD                    // iota = 2, THIRD = SECOND (=iota, =2)
		FOURTH  = "Hello world!" // iota = 3, FOURTH = "Hello world!"
		FIFTH                    // iota = 4, FIFTH = FORTH (="Hello world!")
		SIXTH   = iota           // iota = 5, SIXTH = iota (=5)
		SEVENTH                  // iota = 6, SEVENTH = SIXTH (=iota, =6)
	)

	fmt.Printf("FIRST=%d\n", FIRST)
	fmt.Printf("SECOND=%d\n", SECOND)
	fmt.Printf("THIRD=%d\n", THIRD)
	fmt.Printf("FOURTH=%s\n", FOURTH)
	fmt.Printf("FIFTH=%s\n", FIFTH)
	fmt.Printf("SIXTH=%d\n", SIXTH)
	fmt.Printf("SEVENTH=%d\n", SEVENTH)

	iota2Exponential()
}

func iota2Exponential() {
	const (
		A = 2 << iota
		B = 2 << iota
		C = 2 << iota
		D = 2 << iota
		E = 2 << iota
		F = 2 << iota
	)

	fmt.Println(A, B, C, D, E, F)
}
