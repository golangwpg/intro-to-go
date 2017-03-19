package main

import "fmt"

func main() {
	ageMap := map[string]int{"Mark": 43, "Andrew": 28} // HL

	// update Andrew
	ageMap["Andrew"] = 29 // HL

	// remove Mark
	delete(ageMap, "Mark") // HL

	// add Eric
	ageMap["Eric"] = 34 // HL

	for k, v := range ageMap { // HL
		fmt.Printf("%s is %d years old\n", k, v)
	} // HL
}
