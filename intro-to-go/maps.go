package main

import "fmt"

func main() {

	// create map
	ageMap := map[string]int{ // HL
		"Mark":   43,
		"Andrew": 26,
	} // HL

	// length of map
	fmt.Printf("%d people\n", len(ageMap)) // HL

	// range over map
	for k, v := range ageMap { // HL
		fmt.Printf("%s is %d years old\n", k, v)
	} // HL

}
