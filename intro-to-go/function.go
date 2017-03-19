package main

import "fmt"

const Pi = 3.14

// calculate area of circle
func area(radius float64) float64 { // HL
	return Pi * radius * radius
}

func main() {
	fmt.Println(area(3.0)) // HL
}
