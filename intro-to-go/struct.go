package main

import "fmt"

type Person struct { // HL
	FirstName string // HL
	LastName  string // HL
	Age       int    // HL
} // HL

func main() {

	p := Person{ // HL
		FirstName: "Mark",      // HL
		LastName:  "St.Godard", // HL
		Age:       43,          // HL
	} // HL

	fmt.Printf("%v\n", p)
}
