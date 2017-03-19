package main

import (
	"fmt"

	"github.com/markstgodard/intro/circle"
)

func main() {

	// create a circle
	c := circle.New(3.0) // HL

	// call function Area on circle pointer
	area := c.Area() // HL

	fmt.Printf("area of circle: %f\n", area)
}
