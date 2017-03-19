	// value
	p := Person{FirstName: "Mark", LastName: "St.Godard", Age: 43} // HL

	// pointer to p
	mark = &p // HL

	// dereference pointer p
	fmt.Printf("%s\n", mark.FirstName) // HL
