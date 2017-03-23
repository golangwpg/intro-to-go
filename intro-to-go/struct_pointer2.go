package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	p := &Person{FirstName: "Mark", LastName: "St.Godard", Age: 43} // HL

	fmt.Printf("%s\n", p.FirstName) // HL

	personValAt := *p
	fmt.Printf("%s\n", personValAt.LastName) // HL

	fmt.Printf("%d\n", (*p).Age) // HL
}