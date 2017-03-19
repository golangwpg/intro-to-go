package main

import "fmt"

const PI = 3.14

// calculate area of circle
func area(radius float64) float64 { // HL
	return PI * radius * radius
}

func main() {
	fmt.Println(area(3.0)) // HL
}
